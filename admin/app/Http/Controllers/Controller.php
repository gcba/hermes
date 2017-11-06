<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use TCG\Voyager\Facades\Voyager;
use TCG\Voyager\Http\Controllers\VoyagerBreadController as BreadController;
use Yajra\DataTables\Datatables;
use Validator;

class Controller extends BreadController
{
    // From Voyager's VoyagerBreadController.php, customized

    // POST BRE(A)D
    public function store(Request $request)
    {
        $slug = $this->getSlug($request);
        $dataType = Voyager::model('DataType')->where('slug', '=', $slug)->first();

        // Check permission
        Voyager::canOrFail('add_'.$dataType->name);

        //Validate fields with ajax
        $val = $this->validateBread($request->all(), $dataType->addRows);

        if ($val->fails()) {
            return response()->json(['errors' => $val->messages()]);
        }

        if (!$request->ajax()) {
            $data = $this->insertUpdateData($request, $slug, $dataType->addRows, new $dataType->model_name());

            return redirect()
                ->route("voyager.{$dataType->slug}.index", ['id' => $data->id])
                ->with([
                        'message'    => "Ítem creado exitosamente",
                        'alert-type' => 'success',
                    ]);
        }
    }

    // From Voyager's VoyagerBreadController.php, customized

    // POST BR(E)AD
    public function update(Request $request, $id)
    {
        $slug = $this->getSlug($request);
        $dataType = Voyager::model('DataType')->where('slug', '=', $slug)->first();

        // Check permission
        Voyager::canOrFail('edit_'.$dataType->name);

        //Validate fields with ajax
        $val = $this->validateBread($request->all(), $dataType->editRows);

        if ($val->fails()) {
            return response()->json(['errors' => $val->messages()]);
        }

        if (!$request->ajax()) {
            $data = call_user_func([$dataType->model_name, 'findOrFail'], $id);

            $this->insertUpdateData($request, $slug, $dataType->editRows, $data, $id);

            return redirect()
                ->route("voyager.{$dataType->slug}.index", ['id' => $id])
                ->with([
                    'message'    => "Ítem editado exitosamente",
                    'alert-type' => 'success',
                ]);
        }
    }
}