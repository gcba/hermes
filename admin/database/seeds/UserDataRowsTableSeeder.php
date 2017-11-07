<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\DataRow;
use TCG\Voyager\Models\DataType;

class UserDataRowsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $userDataType = DataType::where('slug', 'users')->firstOrFail();

        $dataRow = $this->dataRow($userDataType, 'id');
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

        $dataRow = $this->dataRow($userDataType, 'name');
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
                'details'      => '{"validation":{"rules":["required","string","min:2","max:70"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 2,
            ])->save();
        }

        $dataRow = $this->dataRow($userDataType, 'email');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Email',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 1,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"validation":{"rules":["required","string","email","unique:users,email"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","email":"El campo :attribute debe ser un email válido."}}}',
                'order'        => 3,
            ])->save();
        }

        $dataRow = $this->dataRow($userDataType, 'avatar');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'image',
                'display_name' => 'Avatar',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 1,
                'edit'         => 1,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"validation":{"rules":["image","dimensions:min_width=100,min_height=100,max_width=300,max_height=300"],"messages":{"image":"El campo :attribute debe ser una imagen.","integer":"El :attribute tener un alto y un ancho ancho entre :min_width y :max_width píxeles."}}}',
                'order'        => 4,
            ])->save();
        }

        $dataRow = $this->dataRow($userDataType, 'role_id');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'user_belongsto_role_relationship',
                'display_name' => 'Rol',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 1,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\Role","table":"roles","type":"belongsTo","column":"role_id","key":"id","label":"display_name","pivot_table":"","pivot":"0"}',
                'order'        => 5,
            ])->save();
        }

        $dataRow = $this->dataRow($userDataType, 'apps');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'user_belongstomany_app_relationship',
                'display_name' => 'Aplicaciones',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 1,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"model":"App\\\App","table":"apps","type":"belongsToMany","column":"id","key":"id","label":"name","pivot_table":"app_user","pivot":"1"}',
                'order'        => 6,
            ])->save();
        }

        $dataRow = $this->dataRow($userDataType, 'updated_at');
        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'timestamp',
                'display_name' => 'Última Modificación',
                'required'     => 0,
                'browse'       => 0,
                'read'         => 0,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '',
                'order'        => 7,
            ])->save();
        }

        $dataRow = $this->dataRow($userDataType, 'updated_by');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'user_belongsto_user_relationship',
                'display_name' => 'Modificado Por',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 1,
                'details'      => '{"model":"App\\\User","table":"users","type":"belongsTo","column":"updated_by","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 8,
            ])->save();
        }

        $dataRow = $this->dataRow($userDataType, 'created_at');

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
                'order'        => 9,
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
