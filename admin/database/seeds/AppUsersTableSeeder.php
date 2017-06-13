<?php

use Illuminate\Database\Seeder;
use App\AppUser;

class AppUsersTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (AppUser::count() == 0) {
            AppUser::create([
                'name'  => 'Mariano Gómez',
                'email' => 'mariano@gomez.com',
                'miba_id' => 100
            ]);

            AppUser::create([
                'name'  => 'Esteban Sosa',
                'email' => 'esteban@sosa.com',
                'miba_id' => 101
            ]);

            AppUser::create([
                'name'  => 'Germán Álvarez',
                'email' => 'german@alvarez.com',
                'miba_id' => 102
            ]);

            AppUser::create([
                'name'  => 'Mariela Domínguez',
                'email' => 'mariela@dominguez.com',
                'miba_id' => 103
            ]);

            AppUser::create([
                'name'  => 'Juliana Pérez',
                'email' => 'juliana@perez.com',
                'miba_id' => 104
            ]);

            AppUser::create([
                'name'  => 'Valentina Echeverría',
                'email' => 'valentina@echeverria.com',
                'miba_id' => 105
            ]);
        }
    }
}
