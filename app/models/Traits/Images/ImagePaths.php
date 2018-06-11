<?php

namespace App\Models\Traits\Images;

use Storage;

trait ImagePaths
{
    /**
     * Override the usual route to an image.
     *
     * @return string
     */
    public function getImageUrlAttribute()
    {
        return Storage::url($this->getImagePath());
    }

    /**
     * Link to the thumbnail for this image.
     *
     * @return string
     */
    public function getThumbnailUrlAttribute()
    {
        return Storage::url($this->getThumbnailPath());
    }

    /**
     * Get a link to the "uploaded" version of this file.
     *
     * @return string
     */
    public function getUploadedUrlAttribute()
    {
        return Storage::temporaryUrl($this->getUploadedPath(), now()->addMinutes(10));
    }

    /**
     * Get the path to the "thumbnail" version of this file.
     *
     * @return string
     */
    public function getImagePath()
    {
        return $this->getImagesFolder() . "/{$this->id}.jpeg";
    }

    /**
     * Get the path to the "thumbnail" version of this file.
     *
     * @return string
     */
    public function getThumbnailPath()
    {
        return $this->getThumbnailsFolder() . "/{$this->id}.jpeg";
    }

    /**
     * Get the path to the "uploaded" version of this file.
     *
     * @return string
     */
    public function getUploadedPath()
    {
        return $this->getUploadsFolder() . "/{$this->uploaded_filename}";
    }

    /**
     * Get the images folder name.
     *
     * @return string
     */
    public function getImagesFolder()
    {
        return defined(static::class . '::IMAGES') ? static::IMAGES : 'images';
    }

    /**
     * Get the thumbnails folder name.
     *
     * @return string
     */
    public function getThumbnailsFolder()
    {
        return defined(static::class . '::THUMBNAILS') ? static::THUMBNAILS : 'thumbnails';
    }

    /**
     * Get the uploads folder name.
     *
     * @return string
     */
    public function getUploadsFolder()
    {
        return defined(static::class . '::UPLOADS') ? static::UPLOADS : 'uploads';
    }
}
