<?php

use Illuminate\Database\Seeder;
use App\Range;

class RangesTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Range::count() == 0) {
            Range::create([
                'from' => 0,
                'to' => 5,
                'key'  => md5("123456")
            ]);

            Range::create([
                'from' => 0,
                'to' => 10,
                'key'  => md5("654321")
            ]);

            Range::create([
                'from' => -5,
                'to' => 5,
                'key'  => md5("987654")
            ]);
        }
    }
}
