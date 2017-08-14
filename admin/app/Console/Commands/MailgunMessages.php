<?php

namespace App\Console\Commands;

use Bogardo\Mailgun\Facades\Mailgun;
use TCG\Voyager\Models\Setting;
use Illuminate\Console\Command;

class MailgunMessages extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'mailgun:send {args*}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Send a test email via Mailgun';

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $args = $this->arguments()['args'];

        if (!$args) {
            $this->error('No arguments passed');

            return;
        }

        if (!$args[0]) {
            $this->error('No message to send');

            return;
        }

        if (!$args[1]) {
            $this->error('No email address');

            return;
        }

        if (!filter_var($args[1], FILTER_VALIDATE_EMAIL)) {
            $this->error('Invalid email address');

            return;
        }

        $this->sendEmail($args[0], $args[1]);
    }

    private function sendEmail(string $text, String $email) {
        $result = Mailgun::raw($text, function ($message) use($email) {
            $message->to($email, env('MAILGUN_SENDER', ''))->subject('Mailgun Test');

            return;
        });

        if ($result->status == 200) {
            $this->info('Message sent successfully');
        }
        else {
            $this->error('Error sending email: ' . $result->message);
        }

        return;
    }
}
