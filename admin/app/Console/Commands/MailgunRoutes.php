<?php

namespace App\Console\Commands;

use Mailgun\Mailgun;
use TCG\Voyager\Models\Setting;
use Illuminate\Console\Command;

class MailgunRoutes extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'mailgun:routes {url?} {--delete} {--list}';

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
        $delete = $this->option('delete');
        $list = $this->option('list');

        if (!$url && !$delete && !$list) {
            $this->error('Missing URL');

            return;
        }

        if (!filter_var($url, FILTER_VALIDATE_URL) && !$delete && !$list) {
            $this->error('Invalid URL');

            return;
        }

        if ($url && ($delete || $list)) {
            $this->error('No URL needed');

            return;
        }

        if ($list) {
            $this->listRoutes();
        }
        else if ($delete) {
            $this->deleteRoutes();
        }
        else {
            $this->checkRoute($url);
        }
    }

    private function listRoutes() {
        $routes = $this->client->get("routes");
        $headers = ["ID", 'Description', 'Expression', "Actions", "Priority", "Created At"];
        $rows = [];

        if (count($routes->http_response_body->items) > 0) {
            foreach ($routes->http_response_body->items as $key => $value) {
                $row = [];

                $row["ID"] = $value->id;
                $row["Description"] = $value->description;
                $row["Expression"] = $value->expression;
                $row["Actions"] = join(', ', $value->actions);
                $row["Priority"] = $value->priority;
                $row["Created At"] = $value->created_at;

                $rows[] = $row;
            }

            $this->table($headers, $rows);
        }
        else {
            $this->info('No routes to show');
        }
    }

    private function deleteRoutes() {
        $routes = $this->client->get("routes");

        if (count($routes->http_response_body->items) > 0) {
            foreach ($routes->http_response_body->items as $key => $value) {
                $res = $this->client->delete('routes/' . $value->id);

                $this->info('Route "' . $value->description . '" deleted successfully');
            }
        }
        else {
            $this->error('No routes to delete');
        }
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
            $id = $newRoute->http_response_body->route->id;
            $description = $newRoute->http_response_body->route->description;

            $this->saveRoute($id, $description);

            return;
        }

        $this->error('Error creating the route. Are you connected to the internet?');
    }

    private function saveRoute(String $id, String $name) {
        $setting = Setting::firstOrNew(['key' => $this->key]);

        $setting->key = $this->key;
        $setting->value = $id;
        $setting->display_name = 'Mailgun Route ID';
        $setting->type = 'text';
        $setting->save();

        $this->info('Route "' . $name . '" saved successfully');
    }
}
