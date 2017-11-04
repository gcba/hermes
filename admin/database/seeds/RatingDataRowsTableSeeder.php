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
                'display_name' => '⭐',
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
                'type'         => 'relationship',
                'field'        => 'rating_belongsto_range_relationship',
                'display_name' => 'Rgo.',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\Range","table":"ranges","type":"belongsTo","column":"range_id","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 3,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'description');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Desc.',
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
                'display_name' => 'Mje.',
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
                'type'         => 'relationship',
                'field'        => 'rating_belongsto_app_relationship',
                'display_name' => 'App',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\App","table":"apps","type":"belongsTo","column":"app_id","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 6,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'app_version');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Vers.',
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
                'type'         => 'relationship',
                'field'        => 'rating_belongsto_platform_relationship',
                'display_name' => 'Plataf.',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\Platform","table":"platforms","type":"belongsTo","column":"platform_id","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 8,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'platform_version');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Vers.',
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

        $dataRow = $this->dataRow($ratingDataType, 'browser_id');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'rating_belongsto_browser_relationship',
                'display_name' => 'Browser',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\Browser","table":"browsers","type":"belongsTo","column":"browser_id","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 10,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'browser_version');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Vers.',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["string","min:1","max:15","nullable"],"messages":{"string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 11,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'appuser_id');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'rating_belongsto_appuser_relationship',
                'display_name' => 'Usuario',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\AppUser","table":"appusers","type":"belongsTo","column":"appuser_id","key":"id","label":"name","pivot_table":"","pivot":"0"}',
                'order'        => 12,
            ])->save();
        }

        $dataRow = $this->dataRow($ratingDataType, 'device_id');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'rating_belongsto_device_relationship',
                'display_name' => 'Disp.',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\Device","table":"devices","type":"belongsTo","column":"device_id","key":"id","label":"name","pivot_table":"","pivot":"0"}',
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