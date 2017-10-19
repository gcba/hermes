<?php

namespace App\Http\Controllers;

use Response;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use TCG\Voyager\Facades\Voyager;
use Yajra\Datatables\Datatables;
use App\Message;

class ApiController extends Controller {
    /**
    * Gets a single message.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function message(Message $message)
    {
        $user = Auth::user();

        if ($user->hasPermission('read_messages')) {
            return Response::json($message);
        }

        return Response::json([], 401);
    }

    /**
    * Gets all the messages that belong to a rating.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function ratingMessages($id)
    {
        $user = Auth::user();

        if ($user->hasPermission('read_messages')) {
            $messages = Message::where('rating_id', $id)->orderBy('created_at', 'asc')->get();

            return Response::json($messages);
        }

        return Response::json([], 401);
    }
}