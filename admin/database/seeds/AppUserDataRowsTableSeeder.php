<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\DataRow;
use TCG\Voyager\Models\DataType;

class AppUserDataRowsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $appuserDataType = DataType::where('slug', 'appusers')->firstOrFail();

        $dataRow = $this->dataRow($appuserDataType, 'id');
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

        $dataRow = $this->dataRow($appuserDataType, 'name');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Nombre',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","string","min:3","max:70"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 2,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'email');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Email',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","string","email","min:3","max:100"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","email":"El campo :attribute debe ser un email válido.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 3,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'ratings');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'select_dropdown',
                'display_name' => 'Calificaciones',
                'required'     => 1,
                'browse'       => 0,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"relationship":{"key":"id","label":"rating","page_slug":"admin/ratings"}}',
                'order'        => 4,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'apps');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'appuser_belongstomany_app_relationship',
                'display_name' => 'Aplicaciones',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"model":"App\\\App","table":"apps","type":"belongsToMany","column":"id","key":"id","label":"name","pivot_table":"app_user_app","pivot":"1"}',
                'order'        => 5,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'platforms');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'appuser_belongstomany_platform_relationship',
                'display_name' => 'Plataformas',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"model":"App\\\Platform","table":"platforms","type":"belongsToMany","column":"id","key":"id","label":"name","pivot_table":"app_user_platform","pivot":"1"}',
                'order'        => 6,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'devices');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'select_multiple',
                'display_name' => 'Dispositivos',
                'required'     => 1,
                'browse'       => 0,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"relationship":{"key":"id","label":"name","page_slug":"admin/devices"}}',
                'order'        => 7,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'miba_id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'ID MiBA',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","string"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto."}}}',
                'order'        => 8,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'updated_at');
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
                'order'        => 9,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'created_at');
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
                'order'        => 10,
            ])->save();
        }

        $dataRow = $this->dataRow($appuserDataType, 'deleted_at');
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
                'order'        => 11,
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