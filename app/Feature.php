<?php

namespace App;

use App\Models\Cacheable;

/**
 * A feature of an Item (e.g. Back Shirring).
 *
 * @property string $name The name of this Feature.
 * @property string $slug The URL slug of this Feature.
 * @property \App\Item[]|\Illuminate\Database\Eloquent\Collection $items
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
     * Hidden attributes on this model.
     *
     * @var array
     */
    protected $hidden = [
        'id',
        'created_at',
        'updated_at',
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
