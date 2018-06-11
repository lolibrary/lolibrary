<?php

use Faker\Generator as Faker;

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

$factory->define(App\Feature::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->domainWord,
        'slug' => str_slug($name),
    ];
});

$factory->define(App\Color::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->colorName,
        'slug' => str_slug($name),
    ];
});

$factory->define(App\Category::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->name,
        'slug' => str_slug($name),
    ];
});

$factory->define(App\Brand::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->name('female'),
        'slug' => str_slug($name),
        'short_name' => str_slug($name),
        'image_id' => uuid4(),
    ];
});

$factory->define(App\Attribute::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->domainWord,
        'slug' => str_slug($name),
    ];
});

$factory->define(App\Tag::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->word,
        'slug' => str_slug($name),
    ];
});
