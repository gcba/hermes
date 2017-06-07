<?php

use Illuminate\Database\Seeder;
use App\Device;
use App\Brand;
use App\Platform;

class DevicesTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Device::count() == 0) {
            $apple = Brand::where('name', 'Apple')->firstOrFail();
            $samsung = Brand::where('name', 'Samsung')->firstOrFail();
            $google = Brand::where('name', 'Google')->firstOrFail();

            $ios = Platform::where('name', 'iOS')->firstOrFail();
            $android = Platform::where('name', 'Android')->firstOrFail();

            Device::create([
                'name'           => 'iPhone 6s',
                'screen_width'   => 750,
                'screen_height'  => 1334,
                'ppi'            => 326,
                'brand_id'       => $apple->id,
                'platform_id'    => $ios->id
            ]);

            Device::create([
                'name'           => 'Galaxy S7',
                'screen_width'   => 1440,
                'screen_height'  => 2560,
                'ppi'            => 557,
                'brand_id'       => $samsung->id,
                'platform_id'    => $android->id
            ]);

            Device::create([
                'name'           => 'Pixel XL',
                'screen_width'   => 1440,
                'screen_height'  => 2560,
                'ppi'            => 534,
                'brand_id'       => $google->id,
                'platform_id'    => $android->id
            ]);

            Device::create([
                'name'           => 'iPhone 7',
                'screen_width'   => 750,
                'screen_height'  => 1334,
                'ppi'            => 326,
                'brand_id'       => $apple->id,
                'platform_id'    => $ios->id
            ]);

            Device::create([
                'name'           => 'Galaxy J7',
                'screen_width'   => 720,
                'screen_height'  => 1280,
                'ppi'            => 267,
                'brand_id'       => $samsung->id,
                'platform_id'    => $android->id
            ]);
        }
    }
}
