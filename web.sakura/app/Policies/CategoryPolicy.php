<?php

namespace App\Policies;

use App\Models\Category;
use App\Models\User;
use Illuminate\Auth\Access\HandlesAuthorization;

class CategoryPolicy
{
    use HandlesAuthorization;

    /**
     * Can a user view available categories?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function viewAny(User $user)
    {
        return $user->junior();
    }

    /**
     * Can a user view a category?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Category $category
     * @return bool
     */
    public function view(User $user, Category $category)
    {
        return $user->junior();
    }

    /**
     * Can a user create a category?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function create(User $user)
    {
        return $user->admin();
    }

    /**
     * Can a user update a category?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Category $category
     * @return bool
     */
    public function update(User $user, Category $category)
    {
        return $user->admin();
    }

    /**
     * Can a user delete a category?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Category $category
     * @return bool
     */
    public function delete(User $user, Category $category)
    {
        if ($category->items()->count() > 0) {
            return false;
        }

        return $user->admin();
    }
}
