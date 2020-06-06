<?php

use Illuminate\Database\Seeder as LaravelSeeder;

class Seeder extends LaravelSeeder
{
    /**
     * A model to use for seeding.
     *
     * @var string
     */
    protected static $model = '';

    /**
     * The content we want to seed.
     *
     * @var array
     */
    protected static $content = [];

    /**
     * A key used for the "value" or "name" column.
     *
     * @var string
     */
    protected static $name = 'name';

    /**
     * The column used for the slug.
     *
     * @var string
     */
    protected static $slug = 'slug';

    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        foreach (static::$content as $slug => $value) {
            if (is_numeric($slug)) {
                // if we have a raw array, slug the value instead.
                $slug = str_slug($value);
            }

            $model = $this->getModel();

            if ($model->newQuery()->where(static::$slug, $slug)->exists()) {
                continue;
            }

            $model->newQuery()->create([
                static::$slug => $slug,
                static::$name => $value,
            ]);
        }
    }

    /**
     * Get the model for this seeder.
     *
     * @return \App\Model
     */
    protected function getModel()
    {
        $model = static::$model;

        if (! class_exists($model)) {
            throw new RuntimeException("Model {$model} not found.");
        }

        return new $model;
    }
}
