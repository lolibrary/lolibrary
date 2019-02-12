<?php

namespace App\Models\Traits;

use App\Models\User;
use Illuminate\Database\Eloquent\Builder as EloquentBuilder;

trait Publishable
{
    /**
     * Publish this item.
     *
     * @param \App\User|null $user
     * @return void
     */
    public function publish(User $user = null)
    {
        $user = $user ?? auth()->user();

        $this->status = static::PUBLISHED;
        $this->publisher()->associate($user);
        $this->published_at = now();
        $this->save();

    }

    /**
     * Make this item a draft.
     *
     * @return void
     */
    public function unpublish()
    {
        $this->status = static::DRAFT;
        $this->save();

        $this->unsearchable();
    }

    /**
     * Return if an item is published or not.
     *
     * @return bool
     */
    public function published(): bool
    {
        return $this->status === static::PUBLISHED;
    }

    /**
     * Return if this item is a draft.
     *
     * @return bool
     */
    public function draft(): bool
    {
        return $this->status === static::DRAFT;
    }

    /**
     * Scope to drafts only.
     *
     * @param \Illuminate\Database\Eloquent\Builder $builder
     * @param bool $draft
     * @return \Illuminate\Database\Eloquent\Builder|\Illuminate\Database\Query\Builder
     */
    public function scopeDrafts(EloquentBuilder $builder, bool $draft = true)
    {
        return $builder->where('status', $draft ? static::DRAFT : static::PUBLISHED);
    }
}
