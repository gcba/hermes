<?php

use Illuminate\Database\Seeder;
use App\App;
use App\Platform;

class PlatformTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Platform::count() == 0 && App::count() == 0) {
            $ios = Platform::create([
                'name' => 'iOS',
                'key'  => md5("123456")
            ]);
            $android = Platform::create([
                'name' => 'Android',
                'key'  => md5("654321")
            ]);
            $windows = Platform::create([
                'name' => 'Windows',
                'key'  => md5("789123")
            ]);

            $denunciaVial = App::create([
                'name' => 'Denuncia Vial',
                'type' => 'M',
                'key'  => md5("123456")
            ]);
            $miBa = App::create([
                'name' => 'Mi BA',
                'type' => 'M',
                'key'  => md5("654321")
            ]);
            $masSimple = App::create([
                'name' => 'Más Simple',
                'type' => 'M',
                'key'  => md5("789123")
            ]);

            $ios.apps()->attach($denunciaVial.id);
            $ios.apps()->attach($miBa.id);
            $ios.apps()->attach($masSimple.id);

            $android.apps()->attach($denunciaVial.id);
            $android.apps()->attach($miBa.id);
            $android.apps()->attach($masSimple.id);
        }
    }
}
