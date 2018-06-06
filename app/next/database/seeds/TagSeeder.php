<?php

use App\Tag;
use Illuminate\Database\Seeder;

class TagSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $tags = json_decode(file_get_contents(storage_path('seeds/tags.json')));

        $records = collect($tags)->map(function (string $tag) {
            return [
                'id' => uuid4(),
                'name' => $tag,
                'slug' => str_slug($tag),
            ];
        })->all();

        Tag::insert($records);
    }
}
