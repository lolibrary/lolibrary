<?php

namespace App\Models;

/**
 * A forum post.
 *
 * @property string $message
 * @property string $user_id The ID of the {@link \App\User user} this post belongs to.
 * @property string $topic_id The ID of the {@link \App\Topic topic} this post belongs to.
 *
 * @property \App\User $user The {@link \App\User user} who owns this post.
 * @property \App\User $topic The {@link \App\Topic topic} this post belongs to.
 */
class Post extends Model
{
    /**
     * @var array
     */
    protected $fillable = ['message'];

    /**
     * The user who created this post.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo
     */
    public function user()
    {
        return $this->belongsTo(User::class);
    }

    /**
     * The topic this post is under.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo
     */
    public function topic()
    {
        return $this->belongsTo(Topic::class);
    }
}
