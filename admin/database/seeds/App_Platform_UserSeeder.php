<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\User;
use App\App;
use App\Platform;
use App\AppUser;

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

            // Let's attach Apps to Platforms

            $ios.apps()->attach($denunciaVial.id);
            $ios.apps()->attach($miBa.id);
            $ios.apps()->attach($masSimple.id);

            $android.apps()->attach($denunciaVial.id);
            $android.apps()->attach($miBa.id);
            $android.apps()->attach($masSimple.id);

            // Let's attach AppUsers to Platforms

            $appuser1 = AppUser::where('email', 'mariano@gomez.com')->firstOrFail();
            $appuser2 = AppUser::where('email', 'esteban@sosa.com')->firstOrFail();
            $appuser3 = AppUser::where('email', 'german@alvarez.com')->firstOrFail();
            $appuser4 = AppUser::where('email', 'mariela@dominguez.com')->firstOrFail();
            $appuser5 = AppUser::where('email', 'juliana@perez.com')->firstOrFail();
            $appuser6 = AppUser::where('email', 'valentina@echeverria.com')->firstOrFail();

            $ios.appusers()->attach($appuser1.id);
            $ios.appusers()->attach($appuser2.id);
            $ios.appusers()->attach($appuser4.id);

            $android.appusers()->attach($appuser3.id);
            $android.appusers()->attach($appuser5.id);
            $android.appusers()->attach($appuser6.id);

            // Let's attach Users to Apps

            $support1 = User::where('email', 'juan@fernandez.com')->firstOrFail();
            $support2 = User::where('email', 'martina@gimenez.com')->firstOrFail();
            $support3 = User::where('email', 'paula@carrizo.com')->firstOrFail();

            $support1.apps()->attach($denunciaVial.id);
            $support2.apps()->attach($miBa.id);
            $support3.apps()->attach($masSimple.id);
        }
    }
}
