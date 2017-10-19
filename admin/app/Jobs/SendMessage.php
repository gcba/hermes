<?php

namespace App\Jobs;

use App\AppUser;
use App\Message;
use App\Jobs\SetMessageTransportId;
use App\Jobs\SetMessageStatus;
use Illuminate\Bus\Queueable;
use Illuminate\Queue\SerializesModels;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;
use Bogardo\Mailgun\Facades\Mailgun;

class SendMessage implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    protected $subject;
    protected $message;
    protected $replyTo;
    protected $user;

    /**
     * The number of seconds the job can run before timing out.
     *
     * @var int
     */
    public $timeout = 5;

    /**
     * Create a new job instance.
     *
     * @return void
     */
    public function __construct(String $subject, Message $message, Message $replyTo, AppUser $user)
    {
        $this->subject = $subject;
        $this->message = $message;
        $this->replyTo = $replyTo;
        $this->user = $user;
    }

    /**
     * Execute the job.
     *
     * @return void
     */
    public function handle()
    {
        $result = Mailgun::raw($this->message->message, function ($message) {
            if ($this->replyTo !== null && isset($this->replyTo->transport_id)) {
                $message->header('In-Reply-To', '<' . $this->replyTo->transport_id . '>');
            }

            $message->to($this->user->email, env('MAILGUN_SENDER', ''))->subject($this->subject);

            return;
        });

        if ($result->status === 200) {
            $id = filter_var(substr(trim($result->id), 1, -1), FILTER_SANITIZE_EMAIL);

            SetMessageTransportId::dispatch($this->message, $id);

            if ($this->replyTo !== null) {
                SetMessageStatus::dispatch($this->replyTo, 2);
            }
        }
        else {
            error_log($result);
            // throw new Exception("Could not send email to Mailgun. Requeuing...");
        }

        return;
    }
}
