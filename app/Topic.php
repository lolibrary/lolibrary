<?php

namespace App;

/**
 * A forum topic.
 *
 * @property \App\User $user
 * @property \App\Post[]|\Illuminate\Database\Eloquent\Collection $posts
 *
 * @property string $user_id
 *
 * @property string $body
 * @property string $teaser
 * @property string $title
 */
class Topic extends Model
{
    /**
     * Get the user for this topic.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo
     */
    public function user()
    {
        return $this->belongsTo(User::class);
    }

    /**
     * Get a collection of posts that are replies to this topic.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany
     */
    public function posts()
    {
        return $this->hasMany(Post::class);
    }

    /**
     * Cut the body of this topic off at the "read more" marker.
     *
     * @return string
     */
    public function getTeaserAttribute()
    {
        if (str_contains($this->body, '~~~~')) {
            return preg_replace('/~~~~(.*)/s', '', $this->body);
        }

        return $this->body;
    }

    /**
     * Topics have singular route names.
     *
     * @return string
     */
    protected function getRouteShowName()
    {
        return 'topic.show';
    }
}
