<?php

use Illuminate\Database\Seeder;
use App\Range;
use App\App;

class RangesTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Range::count() == 0) {
            $denunciaVial = App::where('name', 'Denuncia Vial')->firstOrFail();
            $miBA = App::where('name', 'Mi BA')->firstOrFail();
            $masSimple = App::where('name', 'Más Simple')->firstOrFail();

            Range::create([
                'from' => 0,
                'to' => 5,
                'app_id' => $denunciaVial->id
            ]);

            Range::create([
                'from' => 0,
                'to' => 10,
                'app_id' => $miBA->id
            ]);

            Range::create([
                'from' => -5,
                'to' => 5,
                'app_id' => $masSimple->id
            ]);
        }
    }
}
