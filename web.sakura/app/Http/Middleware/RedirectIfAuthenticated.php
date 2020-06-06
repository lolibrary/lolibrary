<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Support\Facades\Auth;

class RedirectIfAuthenticated
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @param  string|null  $guard
     * @return mixed
     */
    public function handle($request, Closure $next, $guard = null)
    {
        if (Auth::guard($guard)->check()) {
            // conditional: redirect home if regular user, otherwise /library
            if (Auth::guard($guard)->user()->junior()) {
                return redirect('/library');
            }
            
            return redirect('/');
        }

        return $next($request);
    }
}
