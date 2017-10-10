<?php

namespace App\Http\Controllers;

use App\AppUser;
use App\Jobs\SendMessage;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use TCG\Voyager\Facades\Voyager;
use Validator;

class MessagesController extends DataTablesController
{
    // From Voyager's VoyagerBreadController.php, customized

    // POST BRE(A)D
    public function store(Request $request)
    {
        $response = parent::store($request);
        $subject = $request->input('input');
        $text = $request->input('message');
        $userId = $request->input('user');
        $user = DB::table('appusers')->where('miba_id', $userId)->first();

        if (isset($user->email)) {
            SendMessage::dispatch($subject, $text, $user);
        }

        return $response;
    }
}