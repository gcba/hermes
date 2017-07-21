<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use TCG\Voyager\Facades\Voyager;
use Yajra\Datatables\Datatables;

class ServerSideController extends Controller {
    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function ratingsAPI(Request $request)
    {
        return Datatables::of(\App\Rating::query())->make(true);
    }

    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function messagesAPI(Request $request)
    {
        return Datatables::of(\App\Message::query())->make(true);
    }

    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function devicesAPI(Request $request)
    {
        return Datatables::of(\App\Device::query())->make(true);
    }

    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function appusersAPI(Request $request)
    {
        return Datatables::of(\App\AppUser::query())->make(true);
    }

    public function index(Request $request)
    {
        $slug = $this->getSlug($request);
        $dataType = Voyager::model('DataType')->where('slug', '=', $slug)->first();

        // Check permission
        Voyager::canOrFail('browse_'.$dataType->name);

        $model = app($dataType->model_name);
        $dataTypeContent = collect(new $model);
        $isModelTranslatable = is_bread_translatable($model);

        $view = 'voyager::bread.browse';

        if (view()->exists("voyager::$slug.browse")) {
            $view = "voyager::$slug.browse";
        }

        return view($view, compact('dataType', 'dataTypeContent', 'isModelTranslatable'));
    }
}