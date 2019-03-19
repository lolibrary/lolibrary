<?php

namespace Tests\Feature;

use App\Models\Image;
use Illuminate\Http\UploadedFile;
use Illuminate\Support\Facades\Storage;

class ImageTest extends TestCase
{
    public function testImageCreation()
    {
        Storage::fake();

        $file = UploadedFile::fake()->image('foobar.jpeg');

        $image = Image::from($file);

        Storage::disk()->assertExists(config('cdn.image.folder') . '/' . $image->filename);
    }

    public function testImageUrls()
    {
        Storage::fake();

        $file = UploadedFile::fake()->image('foobar.jpeg');

        $image = Image::from($file);

        $this->assertTrue(str_contains($image->getUrlAttribute(), $image->id));
        $this->assertTrue(str_contains($image->getThumbnailUrlAttribute(), $image->id));
    }

    public function testImageUrlEquality()
    {
        Storage::fake();

        $file = UploadedFile::fake()->image('foobar.jpeg');

        $image = Image::from($file);

        $this->assertEquals($image->getUrlAttribute(), $image->url);
        $this->assertEquals($image->getThumbnailUrlAttribute(), $image->thumbnail_url);
    }
}
