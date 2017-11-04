<?php

namespace App\Http\Controllers;

use App\Jobs\SetMessageStatus;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Response;
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

        if ($user !== null) {
            $userApps = $user->apps()->pluck('id')->toArray();

            if ($user->hasPermission('read_messages') && in_array($message->id, $userApps)) {
                return Response::json($message);
            }
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

        if ($user !== null && $user->hasPermission('read_messages')) {
            $userApps = $user->apps()->pluck('id')->toArray();

            $messages = Message::with('rating.app')->where('rating_id', $id)
            ->whereHas('rating.app', function ($query) use($userApps){
                $query->whereIn('id', $userApps);
            })
            ->orderBy('created_at', 'asc')->get();

             foreach ($messages as $item) {
                if ($item->direction === 'in' && $item->status === 0) {
                    SetMessageStatus::dispatch($item, 1);
                }
            }

            return Response::json($messages);
        }

        return Response::json([], 401);
    }
}