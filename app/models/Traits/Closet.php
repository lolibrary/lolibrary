<?php

namespace App\Models\Traits;

use App\Models\Item;

trait Closet
{
    /**
     * The items a user owns.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany|\App\Item[]
     */
    public function closet()
    {
        return $this->belongsToMany(Item::class, 'closet')->withTimestamps();
    }

    /**
     * Update a user's closet and return if we added to it.
     *
     * @param \App\Models\Item $item
     * @return bool
     */
    public function updateCloset(Item $item)
    {
        $result = $this->closet()->toggle($item);

        return count($result['attached']) > 0;
    }

    /**
     * Check if a user has a specific item in their closet.
     *
     * @param \App\Models\Item $item
     * @return bool
     */
    public function owns(Item $item)
    {
        return $this->closet()->where('item_id', $item->id)->exists();
    }
}
