<?php

namespace App;

use App\Models\HasStatus;
use Illuminate\Http\File;
use App\Models\Images\ImagePaths;

/**
 * An item image.
 *
 * @property string $url The location of this image, as a link.
 * @property string $thumbnail_url The location of this image's thumbnail, as a link.
 * @property string|null $uploaded_url The temporary URL to this image.
 * @property string|null $uploaded_filename The filename of the item that was generated to store it for processing.
 * @property int $status The raw status of this image.
 */
class Image extends Model
{
    use ImagePaths, HasStatus;

    /**
     * A status bitmask mapping.
     *
     * @var array
     */
    protected static $statuses = [
        'uploaded'     => 2 ** 0,
        'optimized'    => 2 ** 1,
        'thumbnailled' => 2 ** 2,
    ];

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

        $filename = $file->store($model->getUploadsFolder());

        $model->id = $id ?? uuid4();
        $model->uploaded_filename = $filename;

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

    /**
     * Alias "image_url" to "url".
     *
     * @return string
     */
    public function getUrlAttribute()
    {
        return $this->getImageUrlAttribute();
    }
}
