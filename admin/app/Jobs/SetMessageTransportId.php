<?php

namespace App\Jobs;

use App\Message;
use Illuminate\Bus\Queueable;
use Illuminate\Queue\SerializesModels;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;

class SetMessageTransportId implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    protected $message;
    protected $transportId;

    /**
     * The number of seconds the job can run before timing out.
     *
     * @var int
     */
    public $timeout = 1;

    /**
     * Create a new job instance.
     *
     * @return void
     */
    public function __construct(Message $message, String $transportId)
    {
        $this->message = $message;
        $this->transportId = $transportId;
    }

    /**
     * Execute the job.
     *
     * @return void
     */
    public function handle()
    {
        $this->message->transportId = $this->transportId;

        if (!$this->message->save()) {
            throw new Exception("Could not set email transport id. Requeuing...");
        }
    }
}
