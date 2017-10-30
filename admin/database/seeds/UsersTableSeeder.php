<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Role;
use App\User;

class UsersTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     *
     * @return void
     */
    public function run()
    {
        if (User::count() == 0) {
            $adminRole = Role::where('name', 'admin')->firstOrFail();
            $supportRole = Role::where('name', 'support')->firstOrFail();
            $userRole = Role::where('name', 'user')->firstOrFail();

            User::create([
                'name'           => 'Admin',
                'email'          => 'admin@admin.com',
                'role_id'        => $adminRole->id,
                'avatar'         => 'users/default.png'
            ]);

            User::create([
                'name'           => 'Juan Fernández',
                'email'          => 'juan@fernandez.com',
                'role_id'        => $supportRole->id,
                'avatar'         => 'users/default.png'
            ]);

            User::create([
                'name'           => 'Martina Giménez',
                'email'          => 'martina@gimenez.com',
                'role_id'        => $supportRole->id,
                'avatar'         => 'users/default.png'
            ]);

            User::create([
                'name'           => 'Paula Carrizo',
                'email'          => 'paula@carrizo.com',
                'role_id'        => $supportRole->id,
                'avatar'         => 'users/default.png'
            ]);

            User::create([
                'name'           => 'Miguel Rodríguez',
                'email'          => 'miguel@rodriguez.com',
                'role_id'        => $userRole->id,
                'avatar'         => 'users/default.png'
            ]);

            User::create([
                'name'           => 'Sofía Estévez',
                'email'          => 'sofia@estevez.com',
                'role_id'        => $userRole->id,
                'avatar'         => 'users/default.png'
            ]);

            User::create([
                'name'           => 'Nicolás Uriarte',
                'email'          => 'nicolas@uriarte.com',
                'role_id'        => $userRole->id,
                'avatar'         => 'users/default.png'
            ]);
        }
    }
}
