<?php

Route::get('tags/search', 'TagController@search')->name('tags.search');
Route::resource('tags', 'TagController', ['only' => ['show', 'index']]);
Route::resource('colors', 'ColorController', ['only' => ['show', 'index']]);
Route::resource('brands', 'BrandController', ['only' => ['show', 'index']]);
Route::resource('categories', 'CategoryController', ['only' => ['show', 'index']]);
Route::resource('attributes', 'AttributeController', ['only' => ['show', 'index']]);
Route::resource('features', 'FeatureController', ['only' => ['show', 'index']]);
Route::resource('items', 'ItemController', ['only' => ['show', 'index']]);

Route::post('search', 'SearchController@search')->name('search');
