<?php

namespace Tests\Feature\Models;

use Tests\Feature\TestCase;
use Illuminate\Support\Facades\Storage;
use App\Models\Traits\Images\ImagePaths;

class ImagePathsTest extends TestCase
{
    public function testDefaultImagesFolder()
    {
        $class = new class {
            use ImagePaths;
        };

        $this->assertEquals('images', $class->getImagesFolder());
    }

    public function testDefaultThumbnailsFolder()
    {
        $class = new class {
            use ImagePaths;
        };

        $this->assertEquals('thumbnails', $class->getThumbnailsFolder());
    }

    public function testDefaultUploadsFolder()
    {
        $class = new class {
            use ImagePaths;
        };

        $this->assertEquals('uploads', $class->getUploadsFolder());
    }

    public function testNonDefaultImagesFolder()
    {
        $class = new class {
            use ImagePaths;
            protected const IMAGES = 'something-else-images';
        };

        $this->assertEquals('something-else-images', $class->getImagesFolder());
    }

    public function testNonDefaultThumbnailsFolder()
    {
        $class = new class {
            use ImagePaths;
            protected const THUMBNAILS = 'something-else-thumbnails';
        };

        $this->assertEquals('something-else-thumbnails', $class->getThumbnailsFolder());
    }

    public function testNonDefaultUploadsFolder()
    {
        $class = new class {
            use ImagePaths;
            protected const UPLOADS = 'something-else-uploads';
        };

        $this->assertEquals('something-else-uploads', $class->getUploadsFolder());
    }

    public function testImagePath()
    {
        $class = new class {
            use ImagePaths;
            protected $filename = 'foobar.png';
        };

        $this->assertEquals('images/foobar.png', $class->getImagePath());
    }

    public function testThumbnailPath()
    {
        $class = new class {
            use ImagePaths;
            protected $thumbnail = 'foobar.jpeg';
        };

        $this->assertEquals('thumbnails/foobar.jpeg', $class->getThumbnailPath());
    }

    public function testUploadedPath()
    {
        $class = new class {
            use ImagePaths;
            protected $uploaded_filename = 'foo.png';
        };

        $this->assertEquals('uploads/foo.png', $class->getUploadedPath());
    }

    public function testImageUrl()
    {
        $class = new class {
            use ImagePaths;
            protected $filename = '00000000-0000-0000-0000-000000000000.jpeg';
        };

        $this->assertEquals($this->endpoint() . 'images/00000000-0000-0000-0000-000000000000.jpeg', $class->getImageUrlAttribute());
    }

    public function testThumbnailUrl()
    {
        $class = new class {
            use ImagePaths;
            protected $thumbnail = '00000000-0000-0000-0000-000000000000.png';
        };

        $this->assertEquals($this->endpoint() . 'thumbnails/00000000-0000-0000-0000-000000000000.png', $class->getThumbnailUrlAttribute());
    }

    public function testUploadedUrl()
    {
        $class = new class {
            use ImagePaths;
            protected $uploaded_filename = 'foobar.png';
        };

        $this->assertTrue(str_contains(
            $class->getUploadedUrlAttribute(),
            '/uploads/foobar.png'
        ));
    }

    /**
     * Get the configured minio endpoint in testing.
     *
     * @return string
     */
    protected function endpoint()
    {
        return Storage::url('/');
    }
}
