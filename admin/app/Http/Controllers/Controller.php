<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use TCG\Voyager\Facades\Voyager;
use TCG\Voyager\Http\Controllers\VoyagerBreadController as BreadController;
use TCG\Voyager\Events\BreadImagesDeleted;
use Yajra\DataTables\Datatables;
use Validator;

class Controller extends BreadController
{
    // From Voyager's VoyagerBreadController.php, customized
    /**
     * Delete all images related to a BREAD item.
     *
     * @param \Illuminate\Database\Eloquent\Model $data
     * @param \Illuminate\Database\Eloquent\Model $rows
     *
     * @return void
     */
    public function deleteBreadImages($data, $rows)
    {
        foreach ($rows as $row) {
            if ($data->{$row->field} !== config('voyager.user.default_avatar', 'users/default.png')) {
                $this->deleteFileIfExists($data->{$row->field});

                $options = json_decode($row->details);

                if (isset($options->thumbnails)) {
                    foreach ($options->thumbnails as $thumbnail) {
                        $ext = explode('.', $data->{$row->field});
                        $extension = '.'.$ext[count($ext) - 1];

                        $path = str_replace($extension, '', $data->{$row->field});

                        $thumb_name = $thumbnail->name;

                        $this->deleteFileIfExists($path.'-'.$thumb_name.$extension);
                    }
                }
            }
        }

        if ($rows->count() > 0) {
            event(new BreadImagesDeleted($data, $rows));
        }
    }
}