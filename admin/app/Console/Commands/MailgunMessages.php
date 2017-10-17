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
    protected $signature = 'mailgun:send
        {message : The email body}
        {email : Where to send the message}
        {--subject=Mailgun Test : Email subject}';

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
        $message = $this->argument('message');
        $email = $this->argument('email');
        $subject = $this->option('subject');

        if (!$message && !$email) {
            $this->error('No arguments passed');

            return;
        }

        if (!$message) {
            $this->error('No message to send');

            return;
        }

        if (!$email) {
            $this->error('No email address');

            return;
        }

        if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
            $this->error('Invalid email address');

            return;
        }

        $this->sendEmail($message, $email, $subject);
    }

    private function sendEmail(string $text, String $email, String $subject) {
        $result = Mailgun::raw($text, function ($message) use($email, $subject) {
            $message->to($email, env('MAILGUN_SENDER', ''))->subject($subject);

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
