<?php

namespace App\Models;

use Illuminate\Http\File;
use App\Models\Traits\HasStatus;
use App\Models\Traits\Images\ImagePaths;

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
     * Set the thumbnail folder to a different value.
     *
     * @var string
     */
    protected const THUMBNAILS = 'thumbnail';

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

        // replace 'uploads/' with '' in the filename.
        $filename = str_replace($model->getUploadsFolder() . '/', '', $filename);

        $model->id = $id ?? uuid4();
        $model->filename = $model->id . '.' . $file->extension();
        $model->thumbnail = $model->id . '.jpeg';
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

    /**
     * Get the filename for an image.
     *
     * Overrides {@see \App\Models\Traits\Images\ImagePaths::getRootImagePath}.
     *
     * @return string
     */
    public function getRootImagePath()
    {
        return $this->filename;
    }
}
