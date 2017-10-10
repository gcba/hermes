<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use TCG\Voyager\Facades\Voyager;
use Validator;

class MessagesController extends Controller
{
    // From Voyager's VoyagerBreadController.php, customized

    // POST BRE(A)D
    public function store(Request $request)
    {
        parent::store($request);


    }
}