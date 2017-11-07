<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\DataRow;
use TCG\Voyager\Models\DataType;

class DeviceDataRowsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $deviceDataType = DataType::where('slug', 'devices')->firstOrFail();

        $dataRow = $this->dataRow($deviceDataType, 'id');

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

        $dataRow = $this->dataRow($deviceDataType, 'brand_id');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'device_belongsto_brand_relationship',
                'display_name' => 'Marca',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"model":"App\\\Brand","table":"brands","type":"belongsTo","column":"brand_id","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 2,
            ])->save();
        }

        $dataRow = $this->dataRow($deviceDataType, 'name');

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
                'details'      => '{"validation":{"rules":["required","string","min:1","max:30","unique:devices,name"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 3,
            ])->save();
        }

        $dataRow = $this->dataRow($deviceDataType, 'screen_width');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'number',
                'display_name' => 'Ancho Pantalla',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer","digits_between:3,5"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero.","digits_between":"El campo :attribute debe estar entre :min y :max."}}}',
                'order'        => 4,
            ])->save();
        }

        $dataRow = $this->dataRow($deviceDataType, 'screen_height');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'number',
                'display_name' => 'Altura Pantalla',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer","digits_between:3,5"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero.","digits_between":"El campo :attribute debe estar entre :min y :max."}}}',
                'order'        => 5,
            ])->save();
        }

        $dataRow = $this->dataRow($deviceDataType, 'ppi');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'number',
                'display_name' => 'PPI',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer","digits_between:3,4"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero.","digits_between":"El campo :attribute debe estar entre :min y :max."}}}',
                'order'        => 6,
            ])->save();
        }

        $dataRow = $this->dataRow($deviceDataType, 'platform_id');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'device_belongsto_platform_relationship',
                'display_name' => 'Plataforma',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\Platform","table":"platforms","type":"belongsTo","column":"platform_id","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 7,
            ])->save();
        }

        $dataRow = $this->dataRow($deviceDataType, 'appusers');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'device_belongstomany_appuser_relationship',
                'display_name' => 'Usuarios',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 0,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"model":"App\\\AppUser","table":"appusers","type":"belongsToMany","column":"id","key":"id","label":"name","pivot_table":"app_user_device","pivot":"1"}',
                'order'        => 8,
            ])->save();
        }

        $dataRow = $this->dataRow($deviceDataType, 'updated_at');

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

        $dataRow = $this->dataRow($deviceDataType, 'created_at');

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

        $dataRow = $this->dataRow($deviceDataType, 'deleted_at');

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