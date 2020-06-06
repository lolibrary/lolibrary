<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Auth::routes(['verify' => true]);

// Homepage
Route::get('/', 'HomeController@homepage')->name('home');
Route::get('/search', 'SearchController@index')->name('search');

// User
Route::prefix('profile')->group(function () {
    Route::get('/', 'ProfileController@profile')->name('profile');
    Route::get('closet', 'ProfileController@closet')->name('closet');
    Route::get('wishlist', 'ProfileController@wishlist')->name('wishlist');
});

// blog posts route.
Route::get('blog/{post}', 'BlogController@show')->name('posts.show');

// categories/brands/features etc
Route::group(['namespace' => 'Items\\'], function () {

    $options = ['only' => ['show', 'index']];

    Route::resource('brands', 'BrandController', $options);
    Route::resource('categories', 'CategoryController', $options);
    Route::resource('features', 'FeatureController', $options);
    Route::resource('colors', 'ColorController', $options);
    Route::resource('tags', 'TagController', $options);

    Route::get('items', 'ItemController@index')->name('items.index');
    Route::get('items/{item}', 'ItemController@show')->name('items.show');
});

Route::get('donate', 'DonationController@index')->name('donate');
Route::get('donate/thanks', 'DonationController@thanks')->name('donate.thanks');
Route::get('donate/paypal', 'DonationController@paypal')->name('donate.paypal');
Route::get('donate/patreon', 'DonationController@patreon')->name('donate.patreon');

Route::group(['middleware' => ['auth', 'verified']], function () {
    Route::put('items/{item}/closet', 'Items\\ItemController@closet')->name('items.closet');
    Route::put('items/{item}/wishlist', 'Items\\ItemController@wishlist')->name('items.wishlist');
});
