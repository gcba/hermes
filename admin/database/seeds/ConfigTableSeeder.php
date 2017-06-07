<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\User;
use App\Config;

class ConfigTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Config::count() == 0) { // Falta la app (cómo insertar en many to many?)
            Config::create([
                'name'  => 'Config 1',
                'key'   => 'key1',
                'value' => 'value1'
            ]);

            Config::create([
                'name'  => 'Config 2',
                'key'   => 'key2',
                'value' => 'value2'
            ]);

            Config::create([
                'name'  => 'Config 3',
                'key'   => 'key3',
                'value' => 'value3'
            ]);
        }
    }
}
