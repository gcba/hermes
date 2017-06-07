<?php

use Illuminate\Database\Seeder;
use App\Message;
use App\Rating;

class MessagesTableSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (Message::count() == 0) {
            $ratingsWithMessages = Rating::where('has_message', true)->findOrFail();

            foreach ($ratingsWithMessages as $rating) {
                $message = 'Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ';
                $message . 'ut labore et dolore magna aliqua.';

                Message::create([
                    'message' => $message,
                    'direction' => 'in',
                    'rating_id' => $rating->id
                ]);
            }
        }
    }
}
