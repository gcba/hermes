<?php

use Illuminate\Database\Seeder;
use App\Rating;
use App\AppUser;
use App\App;
use App\Platform;
use App\Device;
use App\Range;

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
            $ranges = Range::all()->toArray();

            $ios = Platform::where('name', 'iOS')->firstOrFail();
            $android = Platform::where('name', 'Android')->firstOrFail();
            $iosDevices = Device::where('platform_id', $ios->id)->get()->toArray();
            $androidDevices = Device::where('platform_id', $android->id)->get()->toArray();

            $ratingOptions = [1, 2, 3, 4, 5];
            $booleanOptions = [true, false];
            $platformOptions = ['iOS', 'Android'];
            $iosVersionOptions = ['8.0', '9.0'];
            $androidVersionOptions = ['5.1', '6.0'];
            $appVersionOptions = ['1.0', '2.0'];

            foreach ($apps as $app) {
                foreach ($appusers as $appuser) {
                    $rating = $ratingOptions[array_rand($ratingOptions)];
                    $description;

                    switch ($rating) {
                        case 1:
                            $description = 'Muy malo';
                            break;
                        case 2:
                            $description = 'Malo';
                            break;
                        case 3:
                            $description = 'Regular';
                            break;
                        case 4:
                            $description = 'Bueno';
                            break;
                        case 5:
                            $description = 'Muy Bueno';
                            break;
                    }

                    $hasMessage = $booleanOptions[array_rand($booleanOptions)];
                    $rangeId = $ranges[array_rand($ranges)]['id'];
                    $platform = $platformOptions[array_rand($platformOptions)];
                    $platformId = $platform == 'iOS' ? $ios->id : $android->id;
                    $platformVersion = $platform == 'iOS' ?
                        $iosVersionOptions[array_rand($iosVersionOptions)] :
                        $androidVersionOptions[array_rand($androidVersionOptions)];
                    $devices = $platform == 'iOS' ? $iosDevices : $androidDevices;
                    $deviceId = $devices[array_rand($devices)]['id'];

                    Rating::create([
                        'rating' => $rating,
                        'description' => $description,
                        'app_version' => $appVersionOptions[array_rand($appVersionOptions)],
                        'platform_version' => $platformVersion,
                        'has_message' => $hasMessage,
                        'app_id' => $app->id,
                        'range_id' => $rangeId,
                        'appuser_id' => $appuser->id,
                        'platform_id' => $platformId,
                        'device_id' => $deviceId
                    ]);
                }
            }
        }
    }
}
