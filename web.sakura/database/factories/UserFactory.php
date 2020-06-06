<?php

use Faker\Generator as Faker;
use Illuminate\Support\Str;

/* @var \Illuminate\Database\Eloquent\Factory $factory */

/*
|--------------------------------------------------------------------------
| Model Factories
|--------------------------------------------------------------------------
|
| This directory should contain each of the model factory definitions for
| your application. Factories provide a convenient way to generate new
| model instances for testing / seeding your application's database.
|
*/

$factory->define(App\Models\User::class, function (Faker $faker) {
    return [
        'id' => uuid4(),
        'name' => $faker->name,
        'username' => $username = $faker->unique()->userName,
        'email' => 'bikeshed+'.$username.'@lolibrary.org',
        'password' => '$2y$10$TKh8H1.PfQx37YgCzwiKb.KjNyWgaHb9cbcoQgdIVFlYg7B77UdFm', // secret
        'remember_token' => Str::random(10),
        'email_verified_at' => now('UTC')->subHour(),
        'banned' => false,
        'level' => App\Models\User::REGULAR,
    ];
});

$factory->state(App\Models\User::class, 'junior', [
    'level' => App\Models\User::JUNIOR_LOLIBRARIAN,
]);

$factory->state(App\Models\User::class, 'lolibrarian', [
    'level' => App\Models\User::LOLIBRARIAN,
]);

$factory->state(App\Models\User::class, 'senior', [
    'level' => App\Models\User::SENIOR_LOLIBRARIAN,
]);

$factory->state(App\Models\User::class, 'admin', [
    'level' => App\Models\User::ADMIN,
]);

$factory->state(App\Models\User::class, 'developer', [
    'level' => App\Models\User::DEVELOPER,
]);

$factory->state(App\Models\User::class, 'banned', [
    'level' => App\Models\User::BANNED,
    'banned' => true,
]);

$factory->state(App\Models\User::class, 'unverified', [
    'email_verified_at' => null,
]);
