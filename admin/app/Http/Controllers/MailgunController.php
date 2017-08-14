<?php

namespace App\Http\Controllers;

use App\Message;

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

        if (filter_var($inReplyTo, FILTER_VALIDATE_EMAIL)) {
            $messageReplied = Message::where('transport_id', $inReplyTo)->first();

            if ($messageReplied) {
                $message = Message::create([
                    'message' => $data['stripped-text'], // Sanitization happens in mutator
                    'direction' => 'in',
                    'notified' => false,
                    'transport_id' => null,
                    'rating_id' => $messageReplied->rating->id
                ]);

                $message->save();
            }
        }

        return Response::json([], 200);
    }
}
