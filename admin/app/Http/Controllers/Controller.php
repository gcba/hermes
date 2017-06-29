<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;
use Illuminate\Support\Str;
use Intervention\Image\Constraint;
use Intervention\Image\Facades\Image;
use TCG\Voyager\Facades\Voyager;
use TCG\Voyager\Traits\AlertsMessages;
use TCG\Voyager\Http\Controllers\VoyagerBreadController as BreadController;
use Validator;
use Log;

class Controller extends BreadController
{
    public function insertUpdateData($request, $slug, $rows, $data, $id = null)
    {
        $multi_select = [];

        /*
         * Prepare Translations and Transform data
         */
        $translations = is_bread_translatable($data)
                        ? $data->prepareTranslations($request)
                        : [];

        foreach ($rows as $row) {
            $options = json_decode($row->details);

            $content = $this->getContentBasedOnType($request, $slug, $row);

            /*
             * merge ex_images and upload images
             */
            if ($row->type == 'multiple_images' && !is_null($content)) {
                if (isset($data->{$row->field})) {
                    $ex_files = json_decode($data->{$row->field}, true);
                    if (!is_null($ex_files)) {
                        $content = json_encode(array_merge($ex_files, json_decode($content)));
                    }
                }
            }

            if (is_null($content)) {
                // Only set the content back to the previous value when there is really now input for this field
                if (is_null($request->input($row->field)) && isset($data->{$row->field})) {
                    $content = $data->{$row->field};
                }
                if ($row->field == 'password') {
                    $content = $data->{$row->field};
                }
            }

            if ($row->type == 'select_multiple' && property_exists($options, 'relationship')) {
                // Only if select_multiple is working with a relationship
                $multi_select[] = ['row' => $row->field, 'content' => $content];
            } else {
                $data->{$row->field} = $content;
            }
        }

        $result = $id ?
            $data->updateOrCreate(['id' => $id], $data->toArray()) :
            $data->updateOrCreate($data->toArray());

        // Save translations
        if (count($translations) > 0) {
            $result->saveTranslations($translations);
        }

        foreach ($multi_select as $sync_data) {
            $result->{$sync_data['row']}()->sync($sync_data['content']);
        }

        return $result;
    }

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
            ->route("voyager.{$dataType->slug}.edit", ['id' => $id])
            ->with([
                'message'    => "Successfully Updated {$dataType->display_name_singular}",
                'alert-type' => 'success',
                ]);
        }
    }
}
