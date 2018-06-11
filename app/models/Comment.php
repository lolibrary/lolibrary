<?php

namespace App\Models;

/**
 * A comment on an Item.
 *
 * @property string $message The text of the comment.
 * @property string $ip_address The IPv4/IPv6 address of the user who commented.
 *
 * @property string $user_id The ID of the {@link \App\User} who commented.
 * @property string $item_id The ID of the {@link \App\Item} this comment is for.
 *
 * @property \App\User $user The {@link \App\User} who commented.
 * @property \App\Item $item The {@link \App\Item} this comment is for.
 *
 * @property \Carbon\Carbon|null $deleted_at
 */
class Comment extends Model
{
    /**
     * The fillable attributes on this model.
     *
     * @var array
     */
    protected $fillable = ['message'];

    /**
     * The user who posted this comment.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo
     */
    public function user()
    {
        return $this->belongsTo(User::class);
    }

    /**
     * The item this comment is for.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo
     */
    public function item()
    {
        return $this->belongsTo(Item::class);
    }
}
