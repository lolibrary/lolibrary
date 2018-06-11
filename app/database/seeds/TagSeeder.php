<?php

class TagSeeder extends Seeder
{
    /**
     * The model to seed.
     *
     * @var string
     */
    protected static $model = App\Models\Tag::class;

    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        static::$content = json_decode(file_get_contents(storage_path('seeds/tags.json')));

        parent::run();
    }
}
