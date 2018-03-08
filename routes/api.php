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

$options = ['only' => ['index', 'show']];

$router->get('tags/search', 'TagController@search')->name('tags.search');
$router->resource('tags', 'TagController', $options);
$router->resource('colors', 'ColorController', $options);
$router->resource('brands', 'BrandController', $options);
$router->resource('categories', 'CategoryController', $options);
$router->resource('attributes', 'AttributeController', $options);
$router->resource('features', 'FeatureController', $options);
$router->resource('items', 'ItemController', $options);

$router->post('search', 'SearchController@search')->name('search');
