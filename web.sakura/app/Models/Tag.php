<?php

namespace App\Models;

/**
 * A search tag.
 *
 * @property string $slug The URL slug of this tag.
 * @property string $name The friendly name of this tag.
 * @property \App\Models\Item[]|\Illuminate\Database\Eloquent\Collection $items
 */
class Tag extends Model
{
    /**
     * The attributes we can fill on this model.
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
     * Get the items that belong to a tag.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany
     */
    public function items()
    {
        return $this->belongsToMany(Item::class);
    }
}
