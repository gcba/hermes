<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\Response;

class MailgunController extends Controller
{
     /**
     * Instantiate a new controller instance.
     *
     * @return void
     */
     public function __construct()
     {
         $this->middleware('log');
     }

     /**
     * Process a Mailgun message.
     *
     * @param  Request  $request
     * @return Response
     */
    public function __invoke(Request $request)
    {


        return Response::make('', 200);
    }
}
