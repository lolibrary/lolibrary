<?php

namespace App\Providers;

use App\Composers;
use Illuminate\Support\ServiceProvider;

class ViewServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        View::composer('component.categories', Composers\Categories::class);
        View::composer('component.brands', Composers\Brands::class);
        View::composer('component.features', Composers\Features::class);
        View::composer('component.tags', Composers\Tags::class);
        View::composer('component.colors', Composers\Colors::class);
        View::composer('component.years', Composers\Years::class);
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
