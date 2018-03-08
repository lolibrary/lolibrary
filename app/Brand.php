<?php

namespace App;

use App\Models\Cacheable;
use App\Traits\ItemCounts;

/**
 * A lolita brand, e.g. Angelic Pretty.
 *
 * @property string $slug The URL slug of this brand.
 * @property string $name The name of this brand.
 * @property string $description An optional description for this brand.
 * @property string $short_name The short name for this brand.
 *
 * @property string $image_id The ID of the {@link \App\Image image} for this brand.
 *
 * @property \App\Image $image The {@link \App\Image image} representing this brand, usually a logo.
 * @property \App\Item[]|\Illuminate\Database\Eloquent\Collection $items
 */
class Brand extends Model
{
    use Cacheable;

    /**
     * The attributes not protected against mass assignment.
     *
     * @var array
     */
    protected $fillable = ['name', 'short_name', 'slug', 'image_id'];

    /**
     * Visible attributes.
     *
     * @var array
     */
    protected $visible = [
        'name',
        'short_name',
        'slug',
        'image',
    ];

    /**
     * Eager loads.
     *
     * @var array
     */
    protected $with = ['image'];

    /**
     * The image for this brand.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo
     */
    public function image()
    {
        return $this->belongsTo(Image::class);
    }
}
