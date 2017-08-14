<?php

// From https://gist.github.com/paulredmond/14523d3bd8062f9ce48cdd1340b3f171

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Response;

/**
 * Validate Mailgun Webhooks
 * @see https://documentation.mailgun.com/user_manual.html#securing-webhooks
 */
class ValidateMailgun
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle($request, Closure $next)
    {
        if (!$request->isMethod('post')) {
            abort(Response::HTTP_FORBIDDEN, 'Only POST requests are allowed.');
        }

        if ($this->verify($request)) {
            return $next($request);
        }

        abort(Response::HTTP_FORBIDDEN);
    }

    /**
     * Build the signature from POST data
     *
     * @see https://documentation.mailgun.com/user_manual.html#securing-webhooks
     * @param  $request The request object
     * @return string
     */
    private function buildSignature($request)
    {
        return hash_hmac(
            'sha256',
            sprintf('%s%s', $request->input('timestamp'), $request->input('token')),
            config('services.mailgun.secret')
        );
    }

    private function verify($request)
    {
        return $this->buildSignature($request) === $request->input('signature');
    }
}