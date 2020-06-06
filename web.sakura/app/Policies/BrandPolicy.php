<?php

namespace App\Policies;

use App\Models\Brand;
use App\Models\User;
use Illuminate\Auth\Access\HandlesAuthorization;

class BrandPolicy
{
    use HandlesAuthorization;

    /**
     * Can a user view available brands?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function viewAny(User $user)
    {
        return $user->junior();
    }

    /**
     * Can a user view a brand?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Brand $brand
     * @return bool
     */
    public function view(User $user, Brand $brand)
    {
        return $user->junior();
    }

    /**
     * Can a user create a brand?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function create(User $user)
    {
        return $user->admin();
    }

    /**
     * Can a user update a brand?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Brand $brand
     * @return bool
     */
    public function update(User $user, Brand $brand)
    {
        return $user->admin();
    }

    /**
     * Can a user delete a brand?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Brand $brand
     * @return bool
     */
    public function delete(User $user, Brand $brand)
    {
        if ($brand->items()->count() > 0) {
            return false;
        }

        return $user->admin();
    }
}
