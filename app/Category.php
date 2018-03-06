<?php

namespace App;

/**
 * A type of item, e.g. JSK.
 *
 * @property string $slug The URL slug of this type.
 * @property string $name The friendly name of this type.
 * @property \App\Item[]|\Illuminate\Database\Eloquent\Collection $items
 */
class Category extends Model
{
    /**
     * Fillable attributes.
     *
     * @var array
     */
    protected $fillable = ['name', 'slug'];

    /**
     * Hidden fields.
     *
     * @var array
     */
    protected $hidden = [
        'id',
        'created_at',
        'updated_at',
    ];
}
