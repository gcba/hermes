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
use Yajra\Datatables\Datatables;
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
        $user = Auth::user();

        if ($user->hasPermission('browse_messages')) {
            $params = $request->query()['columns'];

            $model = Message::with(['rating', 'rating.app', 'rating.platform', 'rating.appuser'])
                ->select([
                    'messages.id',
                    'messages.message',
                    'messages.direction',
                    'messages.status',
                    'messages.created_at',
                    'messages.rating_id'
                ])
                ->where('direction', '=', 'in')
                ->orderBy('created_at', 'desc')
                ->get()
                ->unique('rating_id');

            $datatables = Datatables::of($model)
                ->filter(function ($query) use($params) {
                    $query = $this->filterQuery($query, $params);
                }, true);

            return $datatables->make(true);
        }

        return Response::json([], 401);
    }

    // From Voyager's VoyagerBreadController.php, customized

    // POST BRE(A)D
    public function create(Request $request)
    {
        Voyager::canOrFail('add_messages');

        $created = 201;
        $badRequest = 400;
        $unprocessableEntity = 422;
        $internalServerError = 500;

        $dataType = Voyager::model('DataType')->where('slug', '=', 'messages')->first();
        $validation = $this->validateBread($request->all(), $dataType->addRows);

        if ($validation->fails()) {
            return response()->json(['errors' => $validation->messages(), 'status' => $badRequest]);
        }

        $ratingId = $request->input('rating');
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

        $message = new Message;

        $message->message = $request->input('message');
        $message->direction = 'out';
        $message->status = 0;
        $message->rating()->associate($rating);

        if (!$message->save()) {
            return response()->json(['errors' => 'Could not save new message.', 'status' => $internalServerError]);
        }

        if (isset($user->email)) {
            $subject = env('MAIL_SUBJECT', 'Gracias por tus comentarios');

            SendMessage::dispatch($subject, $message, $replyTo, $user);
        }
        else {
            return response()->json(['errors' => 'User has no email.', 'status' => $unprocessableEntity]);
        }

        return response()->json(['status' => $created, 'message' => $message]);
    }
}