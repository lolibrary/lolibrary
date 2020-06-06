<?php

namespace App\Policies;

use App\Models\{
    Attribute, Color, Feature, Item, User, Tag
};
use Illuminate\Auth\Access\HandlesAuthorization;

class ItemPolicy
{
    use HandlesAuthorization;

    /**
     * Can a user view available items?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function viewAny(User $user)
    {
        return $user->junior();
    }

    /**
     * Can a user view a item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @return bool
     */
    public function view(User $user, Item $item)
    {
        if ($item->user_id !== $user->id) {
            return $user->lolibrarian();
        }

        return $user->junior();
    }

    /**
     * Can a user create an item draft?
     *
     * @param \App\Models\User $user
     * @return bool
     */
    public function create(User $user)
    {
        return $user->junior();
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @return bool
     */
    public function update(User $user, Item $item)
    {
        if ($item->status === Item::PUBLISHED) {
            // lolibrarians can update items they themselves published
            if ($item->publisher_id === $user->id) {
                return $user->lolibrarian();
            }

            return $user->senior();
        }

        // otherwise, this is a draft:
        // users can update their own drafts if junior.
        // users can update other people's drafts if lolibrarian.

        if ($item->user_id === $user->id) {
            return $user->junior();
        }

        return $user->lolibrarian();
    }

    /**
     * Can a user delete an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @return bool
     */
    public function delete(User $user, Item $item)
    {
        if ($item->status === Item::PUBLISHED) {
            // lolibrarian can delete items they themselves published
            if ($item->publisher_id === $user->id) {
                return $user->lolibrarian();
            }

            // senior lolibrarians can delete published items.
            return $user->senior();
        }

        // junior can delete their own drafts.
        if ($item->user_id === $user->id) {
            return $user->junior();
        }



        // only senior can delete drafts from other people.
        // This is just a separate check so it can be changed easily.
        return $user->senior();
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @return bool
     */
    public function publish(User $user, Item $item)
    {
        // must be senior to unpublish, or the original publisher
        if ($item->status === Item::PUBLISHED) {
            if ($item->publisher_id === $user->id) {
                return $user->lolibrarian();
            }

            return $user->senior();
        }

        // otherwise, this is a draft:
        // users can publish their own drafts if lolibrarian.
        if ($item->user_id === $user->id) {
            return $user->lolibrarian();
        }

        // otherwise senior can publish any draft.
        return $user->senior();
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @param \App\Models\Tag $tag
     * @return bool
     */
    public function attachAnyTag(User $user, Item $item)
    {
        return $this->update($user, $item);
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @param \App\Models\Tag $tag
     * @return bool
     */
    public function detachTag(User $user, Item $item, Tag $tag)
    {
        return $this->update($user, $item);
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @param \App\Models\Attribute $attribute
     * @return bool
     */
    public function attachAnyAttribute(User $user, Item $item)
    {
        return $this->update($user, $item);
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @param \App\Models\Attribute $attribute
     * @return bool
     */
    public function detachAttribute(User $user, Item $item, Attribute $attribute)
    {
        return $this->update($user, $item);
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @param \App\Models\Color $color
     * @return bool
     */
    public function attachAnyColor(User $user, Item $item)
    {
        return $this->update($user, $item);
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @param \App\Models\Color $color
     * @return bool
     */
    public function detachColor(User $user, Item $item, Color $color)
    {
        return $this->update($user, $item);
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @param \App\Models\Feature $feature
     * @return bool
     */
    public function attachAnyFeature(User $user, Item $item)
    {
        return $this->update($user, $item);
    }

    /**
     * Can a user update an item?
     *
     * @param \App\Models\User $user
     * @param \App\Models\Item $item
     * @param \App\Models\Feature $feature
     * @return bool
     */
    public function detachFeature(User $user, Item $item, Feature $feature)
    {
        return $this->update($user, $item);
    }
}
