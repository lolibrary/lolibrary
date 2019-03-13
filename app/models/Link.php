<?php

namespace App\Models;

/**
 * @property string $linkable_type
 * @property string $linkable_id
 * @property string $slug
 * @property \App\Models\Model $linkable
 */
class Link extends Model
{
    /**
     * Fillable items.
     *
     * @var array
     */
    protected $fillable = ['linkable_type', 'linkable_id', 'slug'];

    /**
     * @var bool
     */
    public $timestamps = false;

    // link types (can be added to in future)
    const ITEM = Item::class;
    const USER = User::class;

    /**
     * Get a link by slug.
     *
     * @param string $slug
     * @param string $type
     *
     * @return \App\Models\Link
     */
    public static function get(string $slug, string $type = self::ITEM)
    {
        return static::where('linkable_type', '=', $type)->where('slug', '=', $slug)->firstOrFail();
    }

    /**
     * Get the resource we're linking to.
     *
     * @return \Illuminate\Database\Eloquent\Relations\MorphTo
     */
    public function linkable()
    {
        return $this->morphTo();
    }
}
