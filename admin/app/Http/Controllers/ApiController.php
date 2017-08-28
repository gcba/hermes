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
    * Gets messages.
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
}