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
    public function message(Request $request, Message $message)
    {
        if (!$request->ajax()){
            return $this->show($request, $message->id);
        }

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

    // From TCG\Voyager\Http\Controllers\VoyagerBreadController, modified
    public function show(Request $request, $id)
    {
        $dataType = Voyager::model('DataType')->where('slug', '=', 'messages')->first();
        $relationships = $this->getRelationships($dataType);

        if (strlen($dataType->model_name) != 0) {
            $model = app($dataType->model_name);
            $dataTypeContent = call_user_func([$model->with($relationships), 'findOrFail'], $id);
        } else {
            // If Model doest exist, get data from table name
            $dataTypeContent = DB::table($dataType->name)->where('id', $id)->first();
        }

        // Replace relationships' keys for labels and create READ links if a slug is provided.
        $dataTypeContent = $this->resolveRelations($dataTypeContent, $dataType, true);

        // If a column has a relationship associated with it, we do not want to show that field
        $this->removeRelationshipField($dataType, 'read');

        // Check permission
        $this->authorize('read', $dataTypeContent);

        // Check if BREAD is Translatable
        $isModelTranslatable = is_bread_translatable($dataTypeContent);

        $view = 'voyager::bread.read';

        if (view()->exists("voyager::messages.read")) {
            $view = "voyager::$slug.read";
        }

        return Voyager::view($view, compact('dataType', 'dataTypeContent', 'isModelTranslatable'));
    }
}