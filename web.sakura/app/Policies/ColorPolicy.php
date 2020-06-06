<?php

namespace App\Policies;

use App\Models\Color;
use App\Models\User;
use Illuminate\Auth\Access\HandlesAuthorization;

class ColorPolicy
{
    use HandlesAuthorization;

    /**
     * Can a user view available colors?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function viewAny(User $user)
    {
        return $user->junior();
    }

    /**
     * Can a user view a color?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Color $color
     * @return bool
     */
    public function view(User $user, Color $color)
    {
        return $user->junior();
    }

    /**
     * Can a user create a color?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function create(User $user)
    {
        return $user->admin();
    }

    /**
     * Can a user update a color?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Color $color
     * @return bool
     */
    public function update(User $user, Color $color)
    {
        return $user->admin();
    }

    /**
     * Can a user delete a color?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Color $color
     * @return bool
     */
    public function delete(User $user, Color $color)
    {
        if ($color->items()->count() > 0) {
            return false;
        }

        return $user->admin();
    }
}
