<?php

namespace App\Policies;

use App\Models\Tag;
use App\Models\User;
use Illuminate\Auth\Access\HandlesAuthorization;

class TagPolicy
{
    use HandlesAuthorization;

    /**
     * Can a user view available tags?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function viewAny(User $user)
    {
        return $user->junior();
    }

    /**
     * Can a user view a tag?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Tag $tag
     * @return bool
     */
    public function view(User $user, Tag $tag)
    {
        return $user->junior();
    }

    /**
     * Can a user create a tag?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function create(User $user)
    {
        return $user->admin();
    }

    /**
     * Can a user update a tag?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Tag $tag
     * @return bool
     */
    public function update(User $user, Tag $tag)
    {
        return $user->admin();
    }

    /**
     * Can a user delete a tag?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Tag $tag
     * @return bool
     */
    public function delete(User $user, Tag $tag)
    {
        if ($tag->items()->count() > 0) {
            return false;
        }

        return $user->admin();
    }
}
