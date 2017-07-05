<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Role;
use App\User;
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

            // Let's attach Users to Apps

            $supportRole = Role::where('name', 'support')->firstOrFail();
            $supportUsers = User::with('apps')->where('role_id', $supportRole->id)->get();
            $booleanOptions = [true, false];

            foreach ($supportUsers as $support) {
                $isOwner = $booleanOptions[array_rand($booleanOptions)];

                $support->apps()->attach([$denunciaVial->id, $miBA->id, $masSimple->id]);
            }
        }
    }
}
