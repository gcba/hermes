<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\DataRow;
use TCG\Voyager\Models\DataType;

class MessageDataRowsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $messageDataType = DataType::where('slug', 'messages')->firstOrFail();

        $dataRow = $this->dataRow($messageDataType, 'id');

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

        $dataRow = $this->dataRow($messageDataType, 'message');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text_area',
                'display_name' => 'Texto',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"validation":{"rules":["required","string","min:3","max:1000"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}',
                'order'        => 2,
            ])->save();
        }

        $dataRow = $this->dataRow($messageDataType, 'direction');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'text',
                'display_name' => 'Sentido',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '{"validation":{"rules":["required","string","in:in,out"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","in":"El valor del campo :attribute sólo puede ser \'in\' u \'out\'"}}}',
                'order'        => 3,
            ])->save();
        }

        $dataRow = $this->dataRow($messageDataType, 'rating_id');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'relationship',
                'field'        => 'message_belongsto_rating_relationship',
                'display_name' => 'Rating',
                'required'     => 1,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 1,
                'delete'       => 1,
                'details'      => '{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"model":"App\\\Rating","table":"ratings","type":"belongsTo","column":"rating_id","key":"id","label":"rating","pivot_table":"","pivot":"0"}',
                'order'        => 4,
            ])->save();
        }

        $dataRow = $this->dataRow($messageDataType, 'updated_at');

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
                'order'        => 5,
            ])->save();
        }

        $dataRow = $this->dataRow($messageDataType, 'created_at');

        if (!$dataRow->exists) {
            $dataRow->fill([
                'type'         => 'timestamp',
                'display_name' => 'Fecha',
                'required'     => 0,
                'browse'       => 1,
                'read'         => 1,
                'edit'         => 0,
                'add'          => 0,
                'delete'       => 0,
                'details'      => '',
                'order'        => 6,
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