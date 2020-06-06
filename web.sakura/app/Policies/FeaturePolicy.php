<?php

namespace App\Policies;

use App\Models\Feature;
use App\Models\User;
use Illuminate\Auth\Access\HandlesAuthorization;

class FeaturePolicy
{
    use HandlesAuthorization;

    /**
     * Can a user view available features?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function viewAny(User $user)
    {
        return $user->junior();
    }

    /**
     * Can a user view a feature?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Feature $feature
     * @return bool
     */
    public function view(User $user, Feature $feature)
    {
        return $user->junior();
    }

    /**
     * Can a user create a feature?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function create(User $user)
    {
        return $user->admin();
    }

    /**
     * Can a user update a feature?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Feature $feature
     * @return bool
     */
    public function update(User $user, Feature $feature)
    {
        return $user->admin();
    }

    /**
     * Can a user delete a feature?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Feature $feature
     * @return bool
     */
    public function delete(User $user, Feature $feature)
    {
        if ($feature->items()->count() > 0) {
            return false;
        }

        return $user->admin();
    }
}
