<?php

namespace App\Models\Traits;

use App\Models\User;
use App\Models\Item;
use Illuminate\Database\Eloquent\Builder as EloquentBuilder;

trait Publishable
{
    /**
     * Boot this trait, registering model handlers.
     *
     * @return void
     */
    protected static function bootPublishable()
    {
        static::creating(function (Item $model) {
            $model->submitter()->associate(auth()->user());
        });
    }

    /**
     * Publish this item.
     *
     * @param \App\Models\User|null $user
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
    }

    /**
     * Make this item pending.
     *
     * @return void
     */
    public function setPending()
    {
        $this->status = static::PENDING;
        $this->save();
    }

    /**
     * Return if this item is pending.
     *
     * @return bool
     */
    public function isPending(): bool
    {
        return $this->status === static::PENDING;
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
