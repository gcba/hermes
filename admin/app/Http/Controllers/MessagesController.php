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
    public function store(Request $request)
    {
        Voyager::canOrFail('add_messages');

        $created = 201;
        $badRequest = 400;
        $unprocessableEntity = 422;
        $internalServerError = 500;

        $slug = $this->getSlug($request);
        $dataType = Voyager::model('DataType')->where('slug', '=', $slug)->first();
        $validation = $this->validateBread($request->all(), $dataType->addRows);

        if ($validation->fails()) {
            return response()->json(['errors' => $validation->messages(), 'code' => $badRequest]);
        }

        $text = $request->input('message');
        $ratingID = $request->input('rating');
        $rating = Rating::find($ratingID);

        if ($rating === null) {
            return response()->json(['errors' => "Invalid rating.", 'code' => $unprocessableEntity]);
        }

        $message = new Message;

        $message->message = $text;
        $message->direction = 'out';
        $message->rating()->associate($rating);

        if (!$message->save()) {
            return response()->json(['errors' => "Could not save new message.", 'code' => $internalServerError]);
        }

        $subject = $request->input('subject');
        $text = $request->input('message');
        $userId = $request->input('user');
        $user = AppUser::find($userId);

        if ($user === null) {
            return response()->json(['errors' => "Invalid user.", 'code' => $unprocessableEntity]);
        }

        if (isset($user->email)) {
            SendMessage::dispatch($subject, $message, $user);
        }
        else {
            return response()->json(['errors' => "User has no email.", 'code' => $unprocessableEntity]);
        }

        return response()->json(['code' => $created]);
    }
}