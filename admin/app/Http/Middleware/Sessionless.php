<?php

namespace App\Http\Middleware;

use Closure;

class Sessionless
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request $request
     * @param  \Closure $next
     * @return mixed
     */
    public function handle($request, Closure $next)
    {
        config()->set('session.driver', 'array');
        config()->set('cookie.driver', 'array');

        return $next($request);
    }
}