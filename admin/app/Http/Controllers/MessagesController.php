<?php

namespace App\Http\Controllers;

use App\AppUser;
use App\Rating;
use App\Message;
use App\Jobs\SendMessage;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Response;
use TCG\Voyager\Facades\Voyager;
use Yajra\DataTables\Datatables;
use Validator;

class MessagesController extends DataTablesController
{
    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function messagesAPI(Request $request)
    {
        $user = \Auth::user();

        if ($user !== null && $user->hasPermission('browse_messages')) {
            $userApps = $user->apps()->pluck('id')->toArray();
            $params = $request->query()['columns'];
            $search = false;

            foreach ($params as $index => $column) {
                if (strlen(trim($column['search']['value'])) > 0) {
                    $search = true;
                }
            }

            if ($search) {
                $where = $this->parseParams($params);

                $model = Rating::with('messages', 'app', 'appuser', 'platform')
                    ->select('ratings.*')
                    ->where('has_message', true)
                    ->whereHas('app', function ($query) use($userApps) {
                        $query->whereIn('id', $userApps);
                    });

                $ratings = $model->pluck('id')->toArray();
                $messages = Message::whereIn('rating_id', $ratings);

                foreach ($where as $key => $value) {
                    $messages = $messages->where($value[0], $value[1], $value[2]);
                }

                $messages = $messages->latest()->get();
                $model = $model->get();
                $result = collect([]);

                foreach ($model as $key => $value) {
                    $value = collect($value)->forget('messages');
                    $message = collect($messages->where('rating_id', $value['id'])->first());

                    $value->put('messages', $message);

                    $result = $result->push($value);
                }

                $datatables = Datatables::of($result);
            }
            else {
                $model = Rating::with('latestMessage', 'app', 'appuser', 'platform')
                    ->where('has_message', true)
                    ->whereHas('app', function ($query) use($userApps) {
                        $query->whereIn('id', $userApps);
                    })->get();

                $datatables = Datatables::of($model)
                    ->addColumn('messages', function (Rating $rating) {
                        return collect($rating)['latest_message'];
                    })
                    ->removeColumn('latest_message');
            }

            return $datatables->make(true);
        }

        return Response::json([], 401);
    }

    // POST BRE(A)D
    public function create(Request $request)
    {
        Voyager::canOrFail('add_messages');

        $created = 201;
        $badRequest = 400;
        $unprocessableEntity = 422;
        $internalServerError = 500;

        $ratingId = $request->input('rating');
        $messageText = $request->input('message');

        $inputData = ['ratingId' => $ratingId, 'message' => $messageText];
        $rules = ['ratingId' => 'required|integer|min:1', 'message' => 'required|string|max:1500'];
        $validation = Validator::make($inputData, $rules);

        if ($validation->fails()) {
            return response()->json(['errors' => $validation->messages(), 'status' => $badRequest]);
        }

        $rating = Rating::find($ratingId);

        if ($rating === null) {
            return response()->json(['errors' => 'Invalid rating.', 'status' => $unprocessableEntity]);
        }

        $user = AppUser::find($rating->appuser_id);

        if ($user === null) {
            return response()->json(['errors' => 'Invalid user.', 'status' => $unprocessableEntity]);
        }

        $replyTo = Message::where([['rating_id', '=', $rating->id], ['direction', '=', 'in']])
            ->latest()
            ->first();

        if ($replyTo === null) {
            return response()->json(['errors' => 'Message not found.', 'status' => $internalServerError]);
        }

        if (isset($user->email)) {
            $subject = env('MAIL_SUBJECT', 'Example');
            $message = new Message;

            $message->message = $messageText;
            $message->direction = 'out';
            $message->status = 0;
            $message->rating_id = $rating->id;
            $message->createdBy()->associate(Auth::user());

            if (!$message->save()) {
                return response()->json(['errors' => 'Could not save new message.', 'status' => $internalServerError]);
            }

            SendMessage::dispatch($subject, $message, $replyTo, $user);
        }
        else {
            return response()->json(['errors' => 'User has no email.', 'status' => $unprocessableEntity]);
        }

        return response()->json(['status' => $created, 'message' => $message]);
    }
}