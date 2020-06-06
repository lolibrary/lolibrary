<?php

namespace App\Models\Traits;

use RuntimeException;
use App\Models\Item;

trait Sluggable {
    /**
     * Boot this trait and register model listeners.
     * 
     * @return void
     */
    protected static function bootSluggable()
    {
        static::creating(function (Item $model) {
            if ($model->slug !== null) {
                return;
            }

            // Let the exception cause this part to fail here.
            $model->slug = static::createSlug($model);
        });
    }

    /**
     * Get a slug for an item.
     *
     * @param \App\Models\Item $item
     * @return string
     */
    public static function createSlug(Item $item)
    {
        $candidate = $item->brand->short_name . '-' . str_slug($item->english_name);

        if (! static::where('slug', $candidate)->exists()) {
            return $candidate;
        }

        $attempts = -1;

        do {
            if ($attempts > 255) {
                throw new \RuntimeException("Too many items have the slug prefix [{$candidate}]");
            }

            $try = $candidate . '-' . ++$attempts;
        } while (static::where('slug', $try)->exists());

        return $try;
    }
}