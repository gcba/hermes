<?php

namespace App\Console;

use App\Console\Commands\MailgunRoutes;
use App\Console\Commands\MailgunWebhooks;
use App\Console\Commands\MailgunMessages;
use App\Console\Commands\RolesAssign;
use App\Console\Commands\AdminCreate;
use App\Console\Commands\AdminAssign;
use App\Console\Commands\TokenRatings;
use App\Console\Commands\TokenStats;

use Illuminate\Console\Scheduling\Schedule;
use Illuminate\Foundation\Console\Kernel as ConsoleKernel;

class Kernel extends ConsoleKernel
{
    /**
     * The Artisan commands provided by your application.
     *
     * @var array
     */
    protected $commands = [
        MailgunRoutes::class,
        MailgunWebhooks::class,
        MailgunMessages::class,
        RolesAssign::class,
        AdminCreate::class,
        AdminAssign::class,
        TokenRatings::class,
        TokenStats::class
    ];

    /**
     * Define the application's command schedule.
     *
     * @param  \Illuminate\Console\Scheduling\Schedule  $schedule
     * @return void
     */
    protected function schedule(Schedule $schedule)
    {
        // $schedule->command('inspire')
        //          ->hourly();
    }

    /**
     * Register the Closure based commands for the application.
     *
     * @return void
     */
    protected function commands()
    {
        require base_path('routes/console.php');
    }
}
