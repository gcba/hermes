<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\DataRow;
use TCG\Voyager\Models\DataType;

class AppDataRowsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $appDataType = DataType::where('slug', 'apps')->firstOrFail();

        $dataRow = $this->dataRow($appDataType, 'id');

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

        $dataRow = $this->dataRow($appDataType, 'name');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Nombre',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 1,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"validation":{"rules":["required","string","min:3","max:50"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 2,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'type');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'radio_btn',
                'display_name' => 'Tipo',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 1,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"default":"M","options":{"M":"Móvil","W":"Web"},"validation":{"rules":["required","alpha","size:1"],"messages":{"required":"Falta el campo :attribute.","alpha":"El campo :attribute sólo puede constar de una letra.","size":"El campo :attribute sólo puede constar de una letra."}}}',
                'order'        => 3,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'platforms');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'app_belongstomany_platform_relationship',
                'display_name' => 'Plataformas',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 1,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"model":"App\\\Platform","table":"platforms","type":"belongsToMany","column":"id","key":"id","label":"name","pivot_table":"app_platform","pivot":"1"}',
                'order'        => 4,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'appusers');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'app_belongstomany_appuser_relationship',
                'display_name' => 'Usuarios',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 0,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"model":"App\\\AppUser","table":"appusers","type":"belongsToMany","column":"id","key":"id","label":"name","pivot_table":"app_user_app","pivot":"1"}',
                'order'        => 5,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'users');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'app_belongstomany_user_relationship',
                'display_name' => 'Personal',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 1,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"model":"App\\\User","table":"users","type":"belongsToMany","column":"id","key":"id","label":"name","pivot_table":"app_user","pivot":"1"}',
                'order'        => 6,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'key');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Key',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '',
                'order'        => 7,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'updated_at');

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
                'order'        => 8,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'updated_by');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'app_belongsto_user_relationship',
                'display_name' => 'Modificado Por',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 1,
                'details'      => '{"model":"App\\\User","table":"users","type":"belongsTo","column":"updated_by","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 9,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'created_at');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'timestamp',
                'display_name' => 'Creación',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '',
                'order'        => 10,
            ])->save();
        }

        $dataRow = $this->dataRow($appDataType, 'deleted_at');

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