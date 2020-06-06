<?php

namespace App\Models;

use App\Models\Traits\Cacheable;

/**
 * An attribute.
 *
 * @property string $slug The URL route slug of this model.
 * @property string $name The name of this model.
 * @property string $value The value of this attribute's pivot.
 * @property \App\Models\Pivot $pivot A pivot object containing the value of this attribute.
 * @property \App\Models\Item[]|\Illuminate\Database\Eloquent\Collection $items
 */
class Attribute extends Model
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
        'value',
    ];

    /**
     * Attributes to append to the array form.
     *
     * @var array
     */
    protected $appends = ['value'];

    /**
     * A getter for $model->value.
     *
     * @return string|null
     */
    public function getValueAttribute()
    {
        if (! $this->pivot) {
            return;
        }

        return $this->pivot->value;
    }

    /**
     * Get the items that belong to an attribute.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany
     */
    public function items()
    {
        return $this->belongsToMany(Item::class);
    }
}
