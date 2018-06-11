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

$factory->define(App\Models\Feature::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->domainWord,
        'slug' => str_slug($name),
    ];
});

$factory->define(App\Models\Color::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->colorName,
        'slug' => str_slug($name),
    ];
});

$factory->define(App\Models\Category::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->name,
        'slug' => str_slug($name),
    ];
});

$factory->define(App\Models\Brand::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->name('female'),
        'slug' => str_slug($name),
        'short_name' => str_slug($name),
        'image_id' => uuid4(),
    ];
});

$factory->define(App\Models\Attribute::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->domainWord,
        'slug' => str_slug($name),
    ];
});

$factory->define(App\Models\Tag::class, function (Faker $faker) {
    return [
        'name' => $name = $faker->unique()->word,
        'slug' => str_slug($name),
    ];
});
