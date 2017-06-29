<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;
use Illuminate\Support\Facades\DB;
use Log;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        DB::listen(function ($query) {
          //Log::debug($query->sql);
          //Log::debug($query->bindings);
            DB::listen(function ($query) {
                if (strpos($query->sql, 'insert into') !== false) {
                    \Debugbar::info($query->sql);
                    \Debugbar::info($query->bindings);
                }
            });
        });
    }

    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        //
    }
}
