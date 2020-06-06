<?php

namespace App\Models;

/**
 * A Lolita style.
 *
 * @property string $slug The URL slug of this style
 * @property string $name The name of this tag.
 */
class Style extends Model
{
    /**
     * Fillable attributes.
     *
     * @var array
     */
    protected $fillable = ['name', 'slug'];
}
