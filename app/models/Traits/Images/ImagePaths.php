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
        return $this->getImagesFolder() . '/' . $this->getRootImagePath();
    }

    /**
     * Get the path to the "thumbnail" version of this file.
     *
     * @return string
     */
    public function getThumbnailPath()
    {
        return $this->getThumbnailsFolder() . '/' . $this->getRootThumbnailPath();
    }

    /**
     * Get the path to the "uploaded" version of this file.
     *
     * @return string
     */
    public function getUploadedPath()
    {
        return $this->getUploadsFolder() . '/' . $this->getRootUploadedPath();
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

    /**
     * Get the root image filename for URLs.
     *
     * @return string
     */
    public function getRootImagePath()
    {
        return "{$this->filename}";
    }

    /**
     * Get the root image thumbnail filename for URLs.
     *
     * @return string
     */
    public function getRootThumbnailPath()
    {
        return "{$this->thumbnail}";
    }

    /**
     * Get the root image uploaded filename for URLs.
     *
     * @return string
     */
    public function getRootUploadedPath()
    {
        return "{$this->uploaded_filename}";
    }
}
