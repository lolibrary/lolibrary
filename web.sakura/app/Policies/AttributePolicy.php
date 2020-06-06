<?php

namespace App\Policies;

use App\Models\Attribute;
use App\Models\User;
use Illuminate\Auth\Access\HandlesAuthorization;

class AttributePolicy
{
    use HandlesAuthorization;

    /**
     * Can a user view available attributes?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function viewAny(User $user)
    {
        return $user->junior();
    }

    /**
     * Can a user view available attributes?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function view(User $user, Attribute $attribute)
    {
        return $user->junior();
    }

    /**
     * Can a user create a attribute?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function create(User $user)
    {
        return $user->admin();
    }

    /**
     * Can a user update a attribute?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Attribute $attribute
     * @return bool
     */
    public function update(User $user, Attribute $attribute)
    {
        return $user->admin();
    }

    /**
     * Can a user delete a attribute?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Attribute $attribute
     * @return bool
     */
    public function delete(User $user, Attribute $attribute)
    {
        if ($attribute->items()->count() > 0) {
            return false;
        }

        return $user->admin();
    }
}
