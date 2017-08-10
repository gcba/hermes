<?php

namespace App\Providers;

use Mailgun\Mailgun;
use Illuminate\Support\ServiceProvider;

class MailgunServiceProvider extends ServiceProvider
{
    private $client;

    /**
     * Bootstrap the application services.
     *
     * @return void
     */
    public function boot()
    {
        $this->client = new Mailgun(env('MAILGUN_API_KEY', ''));
    }

    /**
     * Register the application services.
     *
     * @return void
     */
    public function register()
    {
        $this->app->singleton(MailgunServiceProvider::class, function ($app) {
            return new MailgunServiceProvider();
        });
    }
}
