<?php

namespace App\Models;

use App\Models\Traits\Cacheable;

/**
 * A feature of an Item (e.g. Back Shirring).
 *
 * @property string $name The name of this Feature.
 * @property string $slug The URL slug of this Feature.
 * @property \App\Models\Item[]|\Illuminate\Database\Eloquent\Collection $items
 */
class Feature extends Model
{
    use Cacheable;

    /**
     * Fillable attributes.
     *
     * @var array
     */
    protected $fillable = ['name', 'slug'];

    /**
     * Visible attributes.
     *
     * @var array
     */
    protected $visible = [
        'name',
        'slug',
        'url',
    ];

    /**
     * Get the items under a feature.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany
     */
    public function items()
    {
        return $this->belongsToMany(Item::class);
    }
}
