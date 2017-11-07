<?php

namespace App\Http\Controllers;

use App\Message;
use App\Jobs\CreateMessage;
use App\Jobs\SetMessageStatus;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Response;

class MailgunController extends Controller
{
     /**
     * Process a Mailgun message.
     *
     * @param  Request  $request
     * @return Response
     */
    public function receive(Request $request)
    {
        $data = [];

        parse_str($request->getContent(), $data);

        $inReplyTo = filter_var(substr(trim($data['In-Reply-To']), 1, -1), FILTER_SANITIZE_EMAIL);
        $messageId = filter_var(substr(trim($data['Message-Id']), 1, -1), FILTER_SANITIZE_EMAIL);

        if (filter_var($inReplyTo, FILTER_VALIDATE_EMAIL) && filter_var($inReplyTo, FILTER_VALIDATE_EMAIL)) {
            $messageReplied = Message::where('transport_id', $inReplyTo)->first();

            if ($messageReplied !== null) {
                $message = $data['stripped-text']; // Sanitization happens in mutator
                $direction = 'in';
                $transportId = $messageId;
                $ratingId = $messageReplied->rating->id;

                CreateMessage::dispatch($message, $direction, $transportId, $ratingId);
                SetMessageStatus::dispatch($messageReplied, 2); // Sent/Received = 0, Notified/Delivered = 1, Replied = 2
            }
        }
        else {
            return Response::json([], 406);
        }

        return Response::json([], 200);
    }

    /**
     * Sets a message as notified.
     *
     * @param  Request  $request
     * @return Response
     */
    public function notify(Request $request)
    {
        $data = [];

        parse_str($request->getContent(), $data);

        $event = filter_var(trim($data['event']), FILTER_SANITIZE_SPECIAL_CHARS);
        $messageId = filter_var(substr(trim($data['Message-Id']), 1, -1), FILTER_SANITIZE_EMAIL);

        if ($event == 'delivered' && filter_var($messageId, FILTER_VALIDATE_EMAIL)) {
            $message = Message::where('transport_id', $messageId)->first();

            if ($message !== null) {
                SetMessageStatus::dispatch($message, 1); // Sent/Received = 0, Notified/Delivered = 1, Replied = 2
            }
        }
        else {
            return Response::json([], 406);
        }

        return Response::json([], 200);
    }
}
