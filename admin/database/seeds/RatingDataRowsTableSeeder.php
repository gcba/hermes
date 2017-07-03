<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\DataRow;
use TCG\Voyager\Models\DataType;

class RatingDataRowsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $ratingDataType = DataType::where('slug', 'ratings')->firstOrFail();

        $dataRow = $this->dataRow($ratingDataType, 'id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'number',
                'display_name' => 'ID',
                'required'     => 1,
                'browse'       => 0,
                'read'         => 0,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '',
                'order'        => 1,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'rating');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'number',
                'display_name' => 'Calificación',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer","min:-127","max:127"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero.","max":"El campo :attribute puede ser hasta :max.","min":"El campo :attribute no debe ser menor a :min."}}}',
                'order'        => 2,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'range_id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'select_dropdown',
                'display_name' => 'Rango',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/ranges"}}',
                'order'        => 3,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'description');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Descripción',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["string","min:1","max:30","nullable"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 4,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'has_message');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'check',
                'display_name' => 'Con Mensaje',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","boolean"],"messages":{"required":"Falta el campo :attribute.","boolean":"El campo :attribute debe ser verdadero o falso."}}}',
                'order'        => 5,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'app_id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'select_dropdown',
                'display_name' => 'Aplicación',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/apps"}}',
                'order'        => 6,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'app_version');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Versión App',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["string","min:1","max:15","nullable"],"messages":{"string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 7,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'platform_id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'select_dropdown',
                'display_name' => 'Plataforma',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/platforms"}}',
                'order'        => 8,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'platform_version');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Versión Plataforma',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["string","min:1","max:15","nullable"],"messages":{"string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 9,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'browser_version');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Versión Browser',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["string","min:1","max:15","nullable"],"messages":{"string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 10,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'appuser_id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'select_dropdown',
                'display_name' => 'Usuario',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/appusers"}}',
                'order'        => 11,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'device_id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'select_dropdown',
                'display_name' => 'Dispositivo',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/devices"}}',
                'order'        => 12,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'browser_id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'select_dropdown',
                'display_name' => 'Browser',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/browsers"}}',
                'order'        => 13,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'updated_at');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'timestamp',
                'display_name' => 'Última Modificación',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '',
                'order'        => 14,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'created_at');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'timestamp',
                'display_name' => 'Creación',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '',
                'order'        => 15,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'deleted_at');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'timestamp',
                'display_name' => 'Borrado',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 0,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '',
                'order'        => 16,
            ])->save();
        }
    }

    /**
     * [dataRow description].
     *
     * @param [type] $type  [description]
     * @param [type] $field [description]
     *
     * @return [type] [description]
     */
    protected function dataRow($type, $field)
    {
        return DataRow::firstOrNew([
                'data_type_id' => $type->id,
                'field'        => $field,
            ]);
    }
}