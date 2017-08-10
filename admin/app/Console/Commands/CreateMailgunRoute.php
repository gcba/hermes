<?php

namespace App\Console\Commands;

use Mailgun\Mailgun;
use TCG\Voyager\Models\Setting;
use Illuminate\Console\Command;

class CreateMailgunRoute extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'mailgun:routes {url}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Create a new Mailgun Route';

    protected $key = 'MAILGUN_ROUTE';

    protected $client;

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();

        $this->client = new Mailgun(env('MAILGUN_API_KEY', ''));
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $url = $this->argument('url');

        if (!$url) {
            $this->error('Missing URL');

            return;
        }

        $this->checkRoute($url);
    }

    private function checkRoute(String $url) {
        if (Setting::where('key', $this->key)->exists()) {
            $route = Setting::where('key', $this->key)->first()->value;
            $results = $this->client->get("routes");

            if ($results->http_response_code == 200) {
                foreach ($results->http_response_body->items as $key => $value) {
                    if ($value->id == $route) {
                        $this->error('Route already exists');

                        return;
                    }
                }
            }
        }

        $this->createRoute($url);
    }

    private function createRoute(String $url) {
        $newRoute = $this->client->post("routes", [
            'priority'    => 0,
            'expression'  => 'catch_all()',
            'action'      => ['forward("' . $url . '")', 'stop()'],
            'description' => 'Forward all messages to Hermes'
        ]);

        if ($newRoute->http_response_code == 200) {
            $this->saveRoute($newRoute->http_response_body->route->id);

            return;
        }

        $this->error('Error creating the route. Are you connected to the internet?');
    }

    private function saveRoute(String $id) {
        $setting = Setting::firstOrNew(['key' => $this->key]);

        $setting->key = $this->key;
        $setting->value = $id;
        $setting->display_name = 'Mailgun Route ID';
        $setting->type = 'text';
        $setting->save();

        $this->info('Route saved: ' . $id);
    }
}
