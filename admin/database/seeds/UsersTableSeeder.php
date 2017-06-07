<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Role;
use TCG\Voyager\Models\User;

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
                'password'       => bcrypt('password'),
                'remember_token' => str_random(60),
                'role_id'        => $adminRole->id
            ]);

            User::create([
                'name'           => 'Juan Fernández',
                'email'          => 'juan@fernandez.com',
                'password'       => bcrypt('password'),
                'remember_token' => str_random(60),
                'role_id'        => $supportRole->id
            ]);

            User::create([
                'name'           => 'Martina Giménez',
                'email'          => 'martina@gimenez.com',
                'password'       => bcrypt('password'),
                'remember_token' => str_random(60),
                'role_id'        => $supportRole->id
            ]);

            User::create([
                'name'           => 'Paula Carrizo',
                'email'          => 'paula@carrizo.com',
                'password'       => bcrypt('password'),
                'remember_token' => str_random(60),
                'role_id'        => $supportRole->id
            ]);

            User::create([
                'name'           => 'Miguel Rodríguez',
                'email'          => 'juan@fernandez.com',
                'password'       => bcrypt('password'),
                'remember_token' => str_random(60),
                'role_id'        => $userRole->id
            ]);

            User::create([
                'name'           => 'Sofía Estévez',
                'email'          => 'sofia@estevez.com',
                'password'       => bcrypt('password'),
                'remember_token' => str_random(60),
                'role_id'        => $userRole->id
            ]);

            User::create([
                'name'           => 'Nicolás Uriarte',
                'email'          => 'nicolas@uriarte.com',
                'password'       => bcrypt('password'),
                'remember_token' => str_random(60),
                'role_id'        => $userRole->id
            ]);
        }
    }
}
