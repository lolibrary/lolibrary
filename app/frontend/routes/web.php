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

// Auth
Auth::routes();
Route::get('auth/verify', 'EmailController@pending')->name('auth.pending');
Route::get('auth/verify/{email}/{token}', 'EmailController@verify')->name('auth.verify');
Route::post('auth/resend', 'EmailController@resend')->name('auth.resend');

// Homepage
Route::get('/', 'HomeController@homepage')->name('home');
Route::get('/search', 'SearchController@index')->name('search');

// User
Route::prefix('profile')->group(function () {
    Route::get('/', 'ProfileController@profile')->name('profile');
    Route::get('closet', 'ProfileController@closet')->name('closet');
    Route::get('wishlist', 'ProfileController@wishlist')->name('wishlist');
});

// Admin
Route::group(['prefix' => 'admin', 'middleware' => ['auth', 'role:junior'], 'as' => 'admin.'], function() {
    Route::get('/', 'AdminController@dashboard')->name('dashboard');
    Route::group(['middleware' => 'role:senior'], function () {
        Route::get('items/queue', 'AdminController@queue')->name('items.queue');
    });
});

// blog posts route.
Route::get('blog/{post}', 'BlogController@show')->name('posts.show');

// categories/brands/features etc
Route::group(['namespace' => 'Items\\'], function () {

    Route::group(['middleware' => ['role:junior']], function () {
        Route::get('items/{item}/edit', 'ItemController@edit')->name('items.edit');
        Route::get('items/create', 'ItemController@create')->name('items.create');
        Route::put('items/store', 'ItemController@store')->name('items.store');
        Route::put('items/{item}/update', 'ItemController@update')->name('items.update');
        Route::delete('items/{item}/image/{image}', 'ItemController@deleteImage')->name('items.images.destroy');
    });

    $options = ['only' => ['show', 'index']];

    Route::resource('brands', 'BrandController', $options);
    Route::resource('categories', 'CategoryController', $options);
    Route::resource('features', 'FeatureController', $options);
    Route::resource('colors', 'ColorController', $options);
    Route::resource('tags', 'TagController', $options);

    Route::get('items', 'ItemController@index')->name('items.index');
    Route::get('items/{item}/show', 'ItemController@show')->name('items.show');



    Route::group(['middleware' => ['role:lolibrarian']], function () {
        Route::put('items/{item}/publish', 'ItemController@publish')->name('items.publish');
        Route::delete('items/{item}/destroy', 'ItemController@destroy')->name('items.destroy');
    });
});

Route::get('donate', 'DonationController@index')->name('donate');
Route::get('donate/thanks', 'DonationController@thanks')->name('donate.thanks');
Route::get('donate/paypal', 'DonationController@paypal')->name('donate.paypal');
Route::get('donate/patreon', 'DonationController@patreon')->name('donate.patreon');

Route::group(['middleware' => ['auth', 'verified']], function () {
    Route::put('items/{item}/closet', 'Items\\ItemController@closet')->name('items.closet');
    Route::put('items/{item}/wishlist', 'Items\\ItemController@wishlist')->name('items.wishlist');
});
