<?php

/** @var \Illuminate\Routing\Router $router */

$router->get('tags/search', 'TagController@search');
$router->get('tags/{tag}', 'TagController@show');
$router->get('tags', 'TagController@index');

$router->get('colors/{color}', 'ColorController@show');
$router->get('colors', 'ColorController@index');

$router->get('brands/{brand}', 'BrandController@show');
$router->get('brands', 'BrandController@index');

$router->get('categories/{category}', 'CategoryController@show');
$router->get('categories', 'CategoryController@index');

$router->get('attributes/{attribute}', 'AttributeController@show');
$router->get('attributes', 'AttributeController@index');

$router->get('features/{feature}', 'FeatureController@show');
$router->get('features', 'FeatureController@index');

$router->get('items/{item}', 'ItemController@show');
$router->get('items', 'ItemController@index');

$router->post('search', 'SearchController@search');
