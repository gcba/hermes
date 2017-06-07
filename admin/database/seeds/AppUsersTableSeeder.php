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
                'email' => 'mariano@gomez.com'
            ]);

            AppUser::create([
                'name'  => 'Esteban Sosa',
                'email' => 'esteban@sosa.com'
            ]);

            AppUser::create([
                'name'  => 'Germán Álvarez',
                'email' => 'german@alvarez.com'
            ]);

            AppUser::create([
                'name'  => 'Mariela Domínguez',
                'email' => 'mariela@dominguez.com'
            ]);

            AppUser::create([
                'name'  => 'Juliana Pérez',
                'email' => 'juliana@perez.com'
            ]);

            AppUser::create([
                'name'  => 'Valentina Echeverría',
                'email' => 'valentina@echeverria.com'
            ]);
        }
    }
}
