<?php

use App\Image;
use Illuminate\Database\Seeder;

class ImageSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        Image::forceCreate([
            'id' => $id = uuid5('default'),
            'filename' => $id . '.png',
            'thumbnail' => $id . '_thumb.png',
            'name' => 'Default Image',
        ]);
    }
}
