<?php

namespace App;

use App\Jobs\DeleteImage;
use Illuminate\Http\File;
use Illuminate\Support\Facades\Storage;

/**
 * An image.
 *
 * @property string $filename This image's filename.
 * @property string $name The filename/description of this image.
 * @property string $thumbnail The thumbnail of this image.
 *
 * @property string $url The location of this image, as a link.
 * @property string $thumbnail_url The location of this image's thumbnail, as a link.
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
        'thumbnail',
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

    protected $hidden = [
        'id',
        'filename',
        'thumbnail',
        'image_id',
        'created_at',
        'updated_at',
    ];

    /**
     * Override the usual route to an image.
     *
     * @return string
     */
    public function getUrlAttribute()
    {
        return Storage::cloud()->url($this->filename);
    }

    /**
     * Link to the thumbnail for this image.
     *
     * @return string
     */
    public function getThumbnailUrlAttribute()
    {
        return Storage::cloud()->url($this->thumbnail);
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
    public static function createFromFile($file, string $id = null)
    {
        $id = $id ?? uuid4();

        $model = new static([
            'name' => $file->getFilename(),
            'filename' => $id . '.' . $file->extension(),
            'thumbnail' => $id . '_thumb.' . $file->extension(),
        ]);

        $model->id = $id;
        $model->save();

        return $model;
    }

    /**
     * Get the default image.
     *
     * @return \App\Image
     * @throws \Illuminate\Database\Eloquent\ModelNotFoundException
     */
    public static function default()
    {
        return static::findOrFail(uuid5('default'));
    }
}
