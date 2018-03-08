<?php

namespace App;

use App\Models\Cacheable;

/**
 * A colorway for an item.
 *
 * @property string $slug The URL slug for this colorway
 * @property string $name The name of this colorway (e.g. Wine)
 * @property \App\Item[]|\Illuminate\Database\Eloquent\Collection $items
 */
class Color extends Model
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
    ];

    /**
     * Get the items that belong to a color.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany
     */
    public function items()
    {
        return $this->belongsToMany(Item::class);
    }
}
