<?php

namespace App\Models;

use Illuminate\Support\Facades\Storage;

/**
 * An item image.
 *
 * @property string $filename               This image's filename
 * @property string $url                    The location of this image, as a link.
 * @property string $thumbnail_url          The location of this image's thumbnail, as a link.
 * @property string|null $uploaded_url      The temporary URL to this image.
 * @property string|null $uploaded_filename The filename of the item that was generated to store it for processing.
 * @property int $status                    The raw status of this image.
 */
class Image extends Model
{
    /**
     * The keys a user is allowed to fill in.
     *
     * @var array
     */
    protected $fillable = [
        'filename',
        'name',
    ];

    /**
     * Extra output to append to this mode's JSON form.
     *
     * @var array
     */
    protected $appends = [
        'url',
        'thumbnail_url',
    ];

    /**
     * Visible attributes.
     *
     * @var array
     */
    protected $visible = [
        'name',
        'url',
        'thumbnail_url',
    ];

    protected static function boot()
    {
        parent::boot();

        // if we can't delete the file, don't delete the image from the database.
        static::deleting(function (self $model) {
            return Storage::disk('cloud')->delete(config('cdn.image.folder') . '/' . $model->filename);
        });
    }

    /**
     * The route key name for images should be its UUID.
     *
     * @return string
     */
    public function getRouteKeyName()
    {
        return $this->primaryKey;
    }

    /**
     * Create an image instance from a file instance.
     *
     * @param \Illuminate\Http\File|\Illuminate\Http\UploadedFile $file
     * @param string $id
     * @return static
     */
    public static function from($file, string $id = null)
    {
        $model = new static;

        $model->id = $id ?? uuid4();
        $model->filename = $model->id . '.' . $file->extension();

        $file->storePubliclyAs(config('cdn.image.folder'), $model->filename);

        $model->save();

        return $model;
    }

    /**
     * Get the default image.
     *
     * @return \App\Models\Image
     * @throws \Illuminate\Database\Eloquent\ModelNotFoundException
     */
    public static function default()
    {
        return static::findOrFail(uuid5('default'));
    }

    /**
     * Create a URL to an image.
     *
     * @return string
     */
    public function getUrlAttribute()
    {
        return cdn_path($this->filename);
    }

    /**
     * Get a URL to a thumbnail of this image.
     *
     * @return string
     */
    public function getThumbnailUrlAttribute()
    {
        return $this->getUrlAttribute() . "?width=300&height=300&fit=bounds";
    }
}
