<?php

namespace App\Models;

class Profile extends Model
{
    /**
     * Fillable attributes.
     *
     * @var array
     */
    protected $fillable = [
        'inspiration',
        'favourite_books',
        'occupation',
        'location',
        'website',
        'since',
        'casual_style',
        'hobbies',
        'livejournal',
    ];

    /**
     * Hidden attributes.
     *
     * @var array
     */
    protected $hidden = [
        'id',
        'url',
        'user_id',
        'user',
        'created_at',
        'updated_at',
    ];

    /**
     * Appended attributes.
     *
     * @var array
     */
    protected $appends = [];
}
