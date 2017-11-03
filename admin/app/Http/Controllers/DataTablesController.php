<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use TCG\Voyager\Facades\Voyager;
use Yajra\Datatables\Datatables;
use App\Rating;
use App\Message;
use App\Device;
use App\AppUser;
use Config;

class DataTablesController extends Controller {
    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function ratingsAPI(Request $request)
    {
        $user = Auth::user();

        if ($user !== null && $user->hasPermission('browse_ratings')) {
            $userApps = $user->apps()->pluck('id')->toArray();
            $params = $request->query()['columns'];
            $where = $this->parseParams($params);

            $model = Rating::with(['range', 'app', 'platform', 'browser', 'appuser', 'device'])
                ->select('ratings.*')
                ->whereHas('app', function ($query) use($userApps) {
                    $query->whereIn('id', $userApps);
                });

            foreach ($where as $key => $value) {
                $model = $model->whereHas($key, function ($query) use($value) {
                    $query->where($value[0], $value[1], $value[2]);
                });
            }

            $datatables = Datatables::of($model)
                ->editColumn('range.name', function($item){
                    return $item->range ? $item->range->name : '';
                })
                ->editColumn('device.name', function($item){
                    return $item->device ? $item->device->name : '';
                })
                ->editColumn('platform.name', function($item){
                    return $item->platform ? $item->platform->name : '';
                })
                ->editColumn('browser.name', function($item){
                    return $item->browser ? $item->browser->name : '';
                })
                ->editColumn('appuser.name', function($item){
                    return $item->appuser ? $item->appuser->name : '';
                });

            return $datatables->make(true);
        }

        return Response::json([], 401);
    }

    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function devicesAPI(Request $request)
    {
        $user = Auth::user();

        if ($user !== null && $user->hasPermission('browse_devices')) {
            $params = $request->query()['columns'];
            $where = $this->parseParams($params);
            $model = Device::with(['platform', 'brand'])->select('devices.*');

            foreach ($where as $key => $value) {
                $model = $model->whereHas($key, function ($query) use($value) {
                    $query->where($value[0], $value[1], $value[2]);
                });
            }

            $datatables = Datatables::of($model)
                ->editColumn('brand.name', function($item) {
                    return $item->brand ? $item->brand->name : '';
                });

            return $datatables->make(true);
        }

        return Response::json([], 401);
    }

    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function appusersAPI(Request $request)
    {
        $user = Auth::user();

        if ($user !== null && $user->hasPermission('browse_appusers')) {
            $userApps = $user->apps()->pluck('id')->toArray();

            $model = AppUser::with('apps')
                ->select('appusers.*')
                ->whereHas('apps', function ($query) use($userApps) {
                    $query->whereIn('id', $userApps);
                });

            return Datatables::of($model)->make(true);
        }

        return Response::json([], 401);
    }

    protected function parseParams($params)
    {
        $where = [];

        foreach ($params as $index => $column) {
            $searchTerm = trim($column['search']['value']);
            $field = explode('.', $column['data']);

            if (strlen($searchTerm) > 0 && count($field) > 1) {
                $isNumeric = is_numeric($searchTerm);
                $operator = $isNumeric ? '=' : 'ilike';
                $searchTerm = $isNumeric ? $searchTerm : '%' . $searchTerm . '%';

                $where[$field[0]] = [$field[1], $operator, $searchTerm];
            }
        }

        return $where;
    }
}