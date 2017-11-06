<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;

class TokenStats extends Token
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'token:stats';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Generate a new token for the Stats API';

    protected function key() {
        return env('HERMES_STATS_PRIVATEKEY', null);
    }
}
