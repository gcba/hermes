<?php

use Illuminate\Database\Seeder;
use App\Browser;

class BrowserTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Browser::count() == 0) {
            Browser::create([
                'name' => 'Chrome'
            ]);

            Browser::create([
                'name' => 'Firefox'
            ]);

            Browser::create([
                'name' => 'Opera'
            ]);
        }
    }
}
