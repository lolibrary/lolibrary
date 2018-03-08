<?php

use Illuminate\Routing\Router;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

/** @var \Illuminate\Routing\Router $router */

$router->group(['domain' => 'api.*'], function (Router $router) {
    $options = ['only' => ['index', 'show']];

    $router->get('tags/search', 'TagController@search')->name('tags.search');
    $router->resource('tags', 'TagController', $options);
    $router->resource('colors', 'ColorController', $options);
    $router->resource('brands', 'BrandController', $options);
    $router->resource('categories', 'CategoryController', $options);
    $router->resource('attributes', 'CategoryController', $options);
});
