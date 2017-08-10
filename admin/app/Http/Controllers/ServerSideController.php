<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use TCG\Voyager\Facades\Voyager;
use Yajra\Datatables\Datatables;
use App\Rating;
use App\Message;
use App\Device;
use App\AppUser;

class ServerSideController extends Controller {
    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function ratingsAPI(Request $request)
    {
        $model = Rating::with(['range', 'app', 'platform', 'browser', 'appuser', 'device'])->select('ratings.*');
        $params = $request->query()['columns'];

        $datatables = Datatables::of($model)
            ->filter(function ($query) use($params) {
                $query = $this->filterQuery($query, $params);
            }, true)
            ->editColumn('appuser.name', function($item){
                return $item->appuser ? $item->appuser->name : '';
            });

        return $datatables->make(true);
    }

    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function messagesAPI(Request $request)
    {
        $model = Message::with('rating')->select('messages.*');
        $params = $request->query()['columns'];

        $datatables = Datatables::of($model)
            ->removeColumn('updated_at')
            ->filter(function ($query) use($params) {
                $query = $this->filterQuery($query, $params);
            }, true)
            ->editColumn('message', function($item){
                return $this->shortenString($item->message, 25);
            });

        return $datatables->make(true);
    }

    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function devicesAPI(Request $request)
    {
        $model = Device::with(['platform', 'brand'])->select('devices.*');
        $params = $request->query()['columns'];

        $datatables = Datatables::of($model)
            ->filter(function ($query) use($params) {
                $query = $this->filterQuery($query, $params);
            }, true)
            ->editColumn('brand.name', function($item){
                    return $item->brand ? $item->brand->name : '';
                });

        return $datatables->make(true);
    }

    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function appusersAPI(Request $request)
    {
        return Datatables::of(AppUser::query())->make(true);
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

    private function filterQuery($query, $params) {
        foreach ($params as $index => $column) {
            $searchTerm = $column['search']['value'];
            $field = explode('.', $column['data']);

            if ($searchTerm !== null && count($field) > 1) {
                $query->whereHas($field[0], function ($q) use ($searchTerm, $field) {
                    $isNumeric = is_numeric($searchTerm);
                    $operator = $isNumeric ? '=' : 'ilike';
                    $searchTerm = $isNumeric ? $searchTerm : $searchTerm . '%';

                    $q->where($field[1], $operator, $searchTerm);
                });
            }
        }

        return $query;
    }

    private function shortenString($string, $limit)
    {
        if(strlen($string) > $limit) {
            $string = trim(substr($string, 0, $limit)) . "...";
        }

        return $string;
    }
}