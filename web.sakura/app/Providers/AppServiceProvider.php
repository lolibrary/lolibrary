<?php

namespace App\Providers;

use App\Composers;
use Illuminate\Support\Facades\View;
use Illuminate\Support\Facades\Blade;
use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        //
    }

    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        View::composer('components.categories', Composers\Categories::class);
        View::composer('components.brands', Composers\Brands::class);
        View::composer('components.features', Composers\Features::class);
        View::composer('components.tags', Composers\Tags::class);
        View::composer('components.colors', Composers\Colors::class);
        View::composer('components.years', Composers\Years::class);

        // Custom 'if' template variables for various roles
        Blade::if('junior', function () {
            return auth()->check() && auth()->user()->junior();
        });
        Blade::if('lolibrarian', function () {
            return auth()->check() && auth()->user()->lolibrarian();
        });
        Blade::if('senior', function () {
            return auth()->check() && auth()->user()->senior();
        });
        Blade::if('admin', function () {
            return auth()->check() && auth()->user()->admin();
        });
        Blade::if('dev', function () {
            return auth()->check() && auth()->user()->developer();
        });
    }
}
