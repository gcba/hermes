<?php

use Illuminate\Database\Seeder;
use App\Brand;

class BrandsTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (\App::isLocal() && Brand::count() == 0) {
            Brand::create([
                'name' => 'Apple'
            ]);

            Brand::create([
                'name' => 'Samsung'
            ]);

            Brand::create([
                'name' => 'Google'
            ]);
        }
    }
}
