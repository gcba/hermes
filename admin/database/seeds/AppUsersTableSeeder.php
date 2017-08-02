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
                'miba_id' => '412319c6-536d-4942-b5bc-0edbd696b9d7'
            ]);

            AppUser::create([
                'name'  => 'Esteban Sosa',
                'email' => 'esteban@sosa.com',
                'miba_id' => '88bf0e47-666b-468d-91b0-b0eb5955298a'
            ]);

            AppUser::create([
                'name'  => 'Germán Álvarez',
                'email' => 'german@alvarez.com',
                'miba_id' => 'dfe21bc0-47ef-40ef-9772-f91c57948510'
            ]);

            AppUser::create([
                'name'  => 'Mariela Domínguez',
                'email' => 'mariela@dominguez.com',
                'miba_id' => '1cb2d373-f265-4e76-a853-a56a1bf30a0e'
            ]);

            AppUser::create([
                'name'  => 'Juliana Pérez',
                'email' => 'juliana@perez.com',
                'miba_id' => '6afb1806-9e44-4622-894e-20538e3a33cf'
            ]);

            AppUser::create([
                'name'  => 'Valentina Echeverría',
                'email' => 'valentina@echeverria.com',
                'miba_id' => 'b8051abe-1a03-4023-9a54-a2ea79c1f8dc'
            ]);
        }
    }
}
