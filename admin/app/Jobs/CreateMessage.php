<?php

namespace App\Jobs;

use App\Message;
use Illuminate\Bus\Queueable;
use Illuminate\Queue\SerializesModels;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;

class CreateMessage implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    protected $message;
    protected $direction;
    protected $status;
    protected $transportId;
    protected $ratingId;

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
    public function __construct(String $message, String $direction, Int $status, $transportId, Int $ratingId)
    {
        $this->message = $message;
        $this->direction = $direction;
        $this->status = $status;
        $this->transportId = $transportId;
        $this->ratingId = $ratingId;
    }

    /**
     * Execute the job.
     *
     * @return void
     */
    public function handle()
    {
        $message = Message::create([
            'message' => $this->message,
            'direction' => $this->direction,
            'status' => $this->status,
            'transport_id' => $this->transportId,
            'rating_id' => $this->ratingId
        ]);

        if (!$message->save()) {
            throw new Exception("Could not save message. Requeuing...");
        }
    }
}
