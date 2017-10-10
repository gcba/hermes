<?php

namespace App\Http\Controllers;

use App\AppUser;
use App\Message;
use App\Jobs\SendMessage;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\DB;
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
            $model = Message::with('rating')->select('messages.*')->where('direction', '=', 'in');
            $params = $request->query()['columns'];

            $datatables = Datatables::of($model)
                ->removeColumn('status')
                ->removeColumn('transport_id')
                ->removeColumn('updated_at')
                ->filter(function ($query) use($params) {
                    $query = $this->filterQuery($query, $params);
                }, true)
                ->editColumn('message', function($item){
                    return $this->shortenString($item->message, 40);
                });

            return $datatables->make(true);
        }

        return Response::json([], 401);
    }

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