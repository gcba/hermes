<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Response;

class MailgunController extends Controller
{
     /**
     * Process a Mailgun message.
     *
     * @param  Request  $request
     * @return Response
     */
    public function receive(Request $request)
    {
        // TODO: Parse response

        return Response::make(null, 200);
    }
}
