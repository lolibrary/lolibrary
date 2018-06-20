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
Route::get('auth/verify/{email}/{token}', 'EmailController@verify')->name('auth.verify');
Route::post('auth/resend', 'EmailController@resend')->name('auth.resend');

// Homepage
Route::get('/', 'HomeController@homepage')->name('home');

// User
Route::prefix('profile')->group(function () {
    Route::get('/', 'ProfileController@profile')->name('profile');
    Route::get('closet', 'ProfileController@closet')->name('closet');
    Route::get('wishlist', 'ProfileController@wishlist')->name('wishlist');
});

// blog posts route.
Route::get('blog/{post}', 'BlogController@show')->name('posts.show');
Route::resource('brands', 'BrandController', ['only' => ['show', 'index']]);
Route::resource('categories', 'CategoryController', ['only' => ['show', 'index']]);
