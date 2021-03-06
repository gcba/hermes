<?php

use Illuminate\Database\Seeder;
use App\User;
use App\Role;
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
                'name' => 'Browser',
                'key'  => md5("789123")
            ]);

            if (\App::isLocal()) {
                $denunciaVial = App::create([
                    'name' => 'Denuncia Vial',
                    'type' => 'M',
                    'key'  => md5("123456")
                ]);

                $miBA = App::create([
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
                $ios->apps()->attach($miBA->id);
                $ios->apps()->attach($masSimple->id);

                $android->apps()->attach($denunciaVial->id);
                $android->apps()->attach($miBA->id);
                $android->apps()->attach($masSimple->id);

                // Let's attach AppUsers to Platforms

                $appusers = AppUser::all();

                $ios->appusers()->attach($appusers[0]->id);
                $ios->appusers()->attach($appusers[1]->id);
                $ios->appusers()->attach($appusers[3]->id);

                $android->appusers()->attach($appusers[2]->id);
                $android->appusers()->attach($appusers[4]->id);
                $android->appusers()->attach($appusers[5]->id);

                // Let's attach AppUsers to Apps

                $denunciaVial->appusers()->attach($appusers[0]->id);
                $denunciaVial->appusers()->attach($appusers[1]->id);
                $miBA->appusers()->attach($appusers[2]->id);
                $miBA->appusers()->attach($appusers[3]->id);
                $masSimple->appusers()->attach($appusers[4]->id);
                $masSimple->appusers()->attach($appusers[5]->id);

                // Let's attach Users to Apps

                $supportRole = Role::where('name', 'support')->firstOrFail();
                $supportUsers = User::where('role_id', $supportRole->id)->get();
                $options = [$denunciaVial->id, $miBA->id, $masSimple->id];

                foreach ($supportUsers as $support) {
                    $support->apps()->attach($options[array_rand($options)]);
                }
            }
        }
    }
}
