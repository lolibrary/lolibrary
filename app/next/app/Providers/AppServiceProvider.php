<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;
use Laravel\Horizon\Horizon;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        Horizon::auth(function () {
            if (! app()->environment('production')) {
                return true;
            }

            return auth()->check() && auth()->user()->developer();
        });
    }

    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        //
    }
}
