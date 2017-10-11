<?php

namespace App\Console\Commands;

use Mailgun\Mailgun;
use TCG\Voyager\Models\Setting;
use Illuminate\Console\Command;

class MailgunWebhooks extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'mailgun:webhooks
    {type? : The type of the webhook to create.}
    {url? : The url of the webhook to create.}
    {--delete : Delete all webhooks. }
    {--list : List all existing webhooks.}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Create a new Mailgun Webhook';
    protected $types = ['open', 'click', 'unsubscribe', 'spam', 'bounce', 'drop', 'deliver'];
    protected $client;
    protected $domain;

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();

        $this->client = new Mailgun(env('MAILGUN_API_KEY', ''));
        $this->domain = env('MAILGUN_DOMAIN', '');
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $type = $this->argument('type');
        $url = $this->argument('url');
        $delete = $this->option('delete');
        $list = $this->option('list');

        if (!$type && !$delete && !$list) {
            $this->error('Missing type');

            return;
        }

        if (!in_array($type, $this->types) && !$delete && !$list) {
            $this->error('Invalid type');

            return;
        }

        if (!$url && !$delete && !$list) {
            $this->error('Missing URL');

            return;
        }

        if (!filter_var($url, FILTER_VALIDATE_URL) && !$delete && !$list) {
            $this->error('Invalid URL');

            return;
        }

        if (($type || $url) && ($delete || $list)) {
            $this->error('No type/URL needed');

            return;
        }

        if ($list) {
            $this->listWebhooks();
        }
        else if ($delete) {
            $this->deleteWebhooks();
        }
        else {
            $this->checkWebhook($type, $url);
        }
    }

    private function listWebhooks() {
        $headers = ['Type', 'URL'];
        $rows = [];
        $webhooks = $this->getWebhooks();

        if ($webhooks !== null) {
            foreach ($this->types as $type) {
                if (array_key_exists($type, $webhooks) && array_key_exists('url', $webhooks[$type])) {
                    $row = [];

                    $row['Type'] = $type;
                    $row['URL'] = $webhooks[$type]['url'];

                    $rows[] = $row;
                }
            }

            $this->table($headers, $rows);
        }
        else {
            $this->info('No webhooks to show');
        }
    }

    private function deleteWebhooks() {
        $webhooks = $this->getWebhooks();

        if ($webhooks === null) {
            foreach ($this->types as $type) {
                if (array_key_exists($type, $webhooks) && array_key_exists('url', $webhooks[$type])) {
                    $response = $this->client->delete("$this->domain/webhooks/$type");

                    $this->info("Webhook '' . $type . '' deleted successfully");
                }
            }
        }
        else {
            $this->error('No webhooks to delete');
        }
    }

    private function checkWebhook(String $type) {
        $webhooks = $this->getWebhooks();

        if ($webhooks !== null && array_key_exists($type, $webhooks) && array_key_exists('url', $webhooks[$type])) {
            $this->error('Webhook already exists');

            return;
        }

        $this->createWebhook($type);
    }

    private function createWebhook(String $type, String $url) {
        $newWebhook = $this->client->post("domains/$this->domain/webhooks", [
            'id'  => $type,
            'url' => $url
        ]);

        if ($newWebhook->http_response_code !== 200) {
            $this->error('Error creating the webhook. Are you connected to the internet?');
        }

        return;
    }

    private function getWebhooks() {
        $response = $this->client->get("domains/$this->domain/webhooks");

        var_dump($response->http_response_body);

        return $response->http_response_code === 200 && isset($response->http_response_body->webhooks) ?
            $response->http_response_body->webhooks :
            null;
    }
}
