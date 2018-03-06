<?php

namespace App;

/**
 * An attribute.
 *
 * @property string $slug The URL route slug of this model.
 * @property string $name The name of this model.
 * @property string $value The value of this attribute's pivot.
 * @property \App\Pivot $pivot A pivot object containing the value of this attribute.
 * @property \App\Item[]|\Illuminate\Database\Eloquent\Collection $items
 */
class Attribute extends Model
{
    /**
     * Fillable attributes.
     *
     * @var array
     */
    protected $fillable = ['name', 'slug'];

    /**
     * Hidden attributes.
     *
     * @var array
     */
    protected $hidden = [
        'id',
        'pivot',
        'created_at',
        'updated_at',
    ];

    /**
     * Attributes to append to the array form.
     *
     * @var array
     */
    protected $appends = ['url', 'value'];

    /**
     * A getter for $model->value.
     *
     * @return string|null
     */
    public function getValueAttribute()
    {
        if (! $this->pivot) {
            return null;
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
