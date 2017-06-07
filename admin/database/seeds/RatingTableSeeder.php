<?php

use Illuminate\Database\Seeder;
use App\Rating;
use App\AppUser;
use App\App;
use App\Platform;
use App\Browser;

class RatingTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Rating::count() == 0) {
            $appuser1 = AppUser::where('email', 'mariano@gomez.com')->firstOrFail();
            $appuser2 = AppUser::where('email', 'esteban@sosa.com')->firstOrFail();
            $appuser3 = AppUser::where('email', 'german@alvarez.com')->firstOrFail();
            $appuser4 = AppUser::where('email', 'mariela@dominguez.com')->firstOrFail();
            $appuser5 = AppUser::where('email', 'juliana@perez.com')->firstOrFail();
            $appuser6 = AppUser::where('email', 'valentina@echeverria.com')->firstOrFail();
            $appusers = [$appuser1, $appuser2, $appuser3, $appuser4, $appuser5, $appuser6, $appuser6];

            $app1 = App::where('name', 'Denuncia Vial')->firstOrFail();
            $app2 = App::where('name', 'Mi BA')->firstOrFail();
            $app3 = AppU::where('name', 'Más Simple')->firstOrFail();
            $apps = [$app1, $app2, $app3];

            $ios = Platform::where('name', 'iOS')->firstOrFail();
            $android = Platform::where('name', 'Android')->firstOrFail();

            Rating::create([
                'rating' => 1,
                'description' => 'Muy Malo',
                'app_version' => '1.0',
                'platform_version' => '8.0',
                'has_message' => false,
                'app_id' => $app1.id,
                'appuser_id' => $appuser1.id,
                'platform_id' => $ios.id,
                'device_id' => $iphone1.id
            ]);

        }
    }
}
