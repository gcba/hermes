<?php

use Illuminate\Database\Seeder;
use App\Rating;
use App\AppUser;
use App\App;
use App\Platform;

class RatingsTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Rating::count() == 0) {
            $appusers = AppUser::all();
            $apps = App::all();

            $ios = Platform::where('name', 'iOS')->firstOrFail();
            $android = Platform::where('name', 'Android')->firstOrFail();

            for ($i = 0; $i < 5; $i++) {
                switch ($i) {
                    case 0:
                        $description = 'Muy malo';
                        break;
                    case 1:
                        $description = 'Malo';
                        break;
                    case 2:
                        $description = 'Regular';
                        break;
                    case 3:
                        $description = 'Bueno';
                        break;
                    case 4:
                        $description = 'Muy Bueno';
                        break;
                }

                 foreach ($apps as $app) {
                    foreach ($appusers as $appuser) {
                        $willLeaveRating = array_rand([true, false]);

                        if ($willLeaveRating) {
                            $description;
                            $deviceId;
                            $hasMessage = array_rand([true, false]);
                            $platform = array_rand(['iOS', 'Android']);
                            $platformId = $platform == "iOS" ? $ios->id : $android->id;
                            $platformVersion = $platform == "iOS" ?
                                array_rand(['8.0', '9.0']) :
                                array_rand(['5.1', '6.0']);
                            $devices = $platform == "iOS" ?
                                array_rand($ios->devices()) :
                                array_rand($android->devices());
                            $deviceId = array_rand($devices)->id;

                            Rating::create([
                                'rating' => $i + 1,
                                'description' => $description,
                                'app_version' => array_rand(['1.0', '2.0']),
                                'platform_version' => $platformVersion,
                                'has_message' => $hasMessage,
                                'app_id' => $app->id,
                                'appuser_id' => $appuser->id,
                                'platform_id' => $platformId,
                                'device_id' => $deviceId
                            ]);
                        }
                    }
                }
            }
        }
    }
}
