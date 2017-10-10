<?php

namespace App\Http\Controllers;

use App\AppUser;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use TCG\Voyager\Facades\Voyager;
use Validator;

class MessagesController extends Controller
{
    // From Voyager's VoyagerBreadController.php, customized

    // POST BRE(A)D
    public function store(Request $request)
    {
        $response = parent::store($request);
        $text = $request->input('message');
        $userId = $request->input('user');
        $subject = $request->input('input');
        $user = DB::table('appusers')->where('miba_id', $userId)->first();

        if (isset($user->email)) {
            $result = Mailgun::raw($text, function ($message, $subject) use($user) {
                $message->to($user->email, env('MAILGUN_SENDER', ''))->subject($subject);

                return;
            });
        }

        return $response;
    }
}