<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;

class TokenRatings extends Token
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'token:ratings';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Generate a new token for the Ratings API';

    protected function key() {
        return env('HERMES_RATINGS_PRIVATEKEY', null);
    }
}