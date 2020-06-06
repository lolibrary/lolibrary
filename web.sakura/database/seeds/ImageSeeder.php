<?php

use App\Models\Image;
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
        $id = uuid5('default');

        if (Image::where('id', $id)->exists()) {
            return;
        }

        Image::forceCreate([
            'id' => $id,
            'filename' => $id . '.png',
            'name' => 'Default Image',
        ]);
    }
}
