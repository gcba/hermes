<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Role;
use App\App;
use App\Platform;
use App\AppUser;

class Apps_Platforms_UsersSeeder extends Seeder
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

            $ios->apps()->attach($denunciaVial->id);
            $ios->apps()->attach($miBa->id);
            $ios->apps()->attach($masSimple->id);

            $android->apps()->attach($denunciaVial->id);
            $android->apps()->attach($miBa->id);
            $android->apps()->attach($masSimple->id);

            // Let's attach AppUsers to Platforms

            $appusers = AppUser::all();

            $ios->appusers()->attach($appuser[0]->id);
            $ios->appusers()->attach($appuser[1]->id);
            $ios->appusers()->attach($appuser[3]->id);

            $android->appusers()->attach($appuser[2]->id);
            $android->appusers()->attach($appuser[4]->id);
            $android->appusers()->attach($appuser[5]->id);

            // Let's attach Users to Apps

            $supportUsers = Role::where('name', 'support')->firstOrFail()->users();

            $supportUsers[0]->apps()->attach($denunciaVial->id);
            $supportUsers[1]->apps()->attach($miBa->id);
            $supportUsers[2]->apps()->attach($masSimple->id);
        }
    }
}
