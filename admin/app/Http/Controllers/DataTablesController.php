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

class DataTablesController extends Controller {
    /**
    * Process datatables ajax request.
    *
    * @return \Illuminate\Http\JsonResponse
    */
    public function ratingsAPI(Request $request)
    {
        $user = Auth::user();

        if ($user->hasPermission('browse_ratings')) {
            $userApps = $user->apps()->pluck('id')->toArray();

            $model = Rating::with(['range', 'app', 'platform', 'browser', 'appuser', 'device'])
                ->select('ratings.*')
                ->whereHas('app', function ($query) use($userApps) {
                    $query->whereIn('id', $userApps);
                });

            $params = $request->query()['columns'];
            $noResults = 'No results';

            $datatables = Datatables::of($model)
                ->filter(function ($query) use($params) {
                    $query = $this->filterQuery($query, $params);
                }, true)
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

        if ($user->hasPermission('browse_devices')) {
            $model = Device::with(['platform', 'brand'])->select('devices.*');
            $params = $request->query()['columns'];

            $datatables = Datatables::of($model)
                ->filter(function ($query) use($params) {
                    $query = $this->filterQuery($query, $params);
                }, true)
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

        if ($user->hasPermission('browse_appusers')) {
            $userApps = $user->apps()->pluck('id')->toArray();

            $model = AppUser::with('apps')
                ->select('appusers.*')
                ->whereHas('apps', function ($query) {
                    $query->whereIn('id', $userApps);
                });

            return Datatables::of($model)->make(true);
        }

        return Response::json([], 401);
    }

    protected function filterQuery($query, $params)
    {
        foreach ($params as $index => $column) {
            $searchTerm = $column['search']['value'];
            $field = explode('.', $column['data']);

            if ($searchTerm !== null && count($field) > 1) {
                $query->whereHas($field[0], function ($q) use ($searchTerm, $field) {
                    $isNumeric = is_numeric($searchTerm);
                    $operator = $isNumeric ? '=' : 'like';
                    $searchTerm = $isNumeric ? $searchTerm : $searchTerm . '%';

                    $q->where($field[1], $operator, $searchTerm);
                });
            }
        }

        return $query;
    }

    protected function shortenString($string, $limit)
    {
        return strlen($string) > $limit ? trim(substr($string, 0, $limit)) . "..." : $string;
    }
}