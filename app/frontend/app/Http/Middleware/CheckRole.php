<?php
namespace App\Http\Middleware;
use Closure;
use Illuminate\Contracts\Auth\Guard;
class CheckRole
{
    /**
     * The guard instance for this request.
     *
     * @var \Illuminate\Contracts\Auth\Guard
     */
    protected $guard;
    /**
     * User for the current request, if any.
     *
     * @var \App\User
     */
    protected $user;
    /**
     * CheckRole constructor.
     *
     * @param \Illuminate\Contracts\Auth\Guard $guard
     */
    public function __construct(Guard $guard)
    {
        $this->guard = $guard;
        $this->user = $guard->user();
    }
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request $request
     * @param  \Closure $next
     * @param string $role
     * @return mixed
     */
    public function handle($request, Closure $next, string $role)
    {
        $allowed = ($this->user && $this->user->{$role}());
        if (! $allowed) {
            return redirect('/')->withErrors("You're not allowed to take that action");
        }
        return $next($request);
    }
}