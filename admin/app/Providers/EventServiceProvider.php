<?php

namespace App\Providers;

use App\App;
use App\Services\UtilsService;

use Illuminate\Support\Facades\Event;
use Illuminate\Foundation\Support\Providers\EventServiceProvider as ServiceProvider;

class EventServiceProvider extends ServiceProvider
{
    /**
     * The event listener mappings for the application.
     *
     * @var array
     */
    protected $listen = [
        'App\Events\Event' => [
            'App\Listeners\EventListener',
        ],
    ];

    /**
     * Register any events for your application.
     *
     * @return void
     */
    public function boot()
    {
        parent::boot();

        Event::listen('eloquent.*taching: App\App', function ($eventName, array $eventData) {
            $app = $eventData['parent_model']::find($eventData['parent_id']);
            $user = \Auth::user();

            if ($app && $user !== null) {
                $app->updated_by = $user->id;

                $app->save();
            }
        });

        Event::listen('eloquent.*taching: App\User', function ($eventName, array $eventData) {
            $parent = $eventData['parent_model']::find($eventData['parent_id']);
            $user = \Auth::user();

            if ($parent && $user !== null) {
                $parent->updated_by = $user->id;

                $parent->save();
            }
        });
    }
}
