<?php

namespace Tests\Feature\Policies;

use App\Models\Item;
use App\Models\User;
use Tests\TestCase;

class ItemPolicyTest extends TestCase
{
    public function test_junior_users_can_delete_their_drafts()
    {
        $user = factory(User::class)->state('junior')->make();
        $item = factory(Item::class)->state('draft')->make(['user_id' => $user->id]);

        $this->assertTrue($user->can('delete', $item));
    }

    public function test_junior_users_cannot_delete_other_drafts()
    {
        $user = factory(User::class)->state('junior')->make();
        $item = factory(Item::class)->state('draft')->make(['user_id' => uuid4()]);

        $this->assertFalse($user->can('delete', $item));
    }

    public function test_junior_users_cannot_delete_published_items()
    {
        $user = factory(User::class)->state('junior')->make();

        $item1 = factory(Item::class)->state('published')->make(['user_id' => $user->id]);
        $item2 = factory(Item::class)->state('published')->make(['user_id' => uuid4()]);

        $this->assertFalse($user->can('delete', $item1));
        $this->assertFalse($user->can('delete', $item2));
    }

    public function test_regular_users_cannot_delete_items()
    {
        $user = factory(User::class)->make();

        $item1 = factory(Item::class)->state('published')->make(['user_id' => $user->id]);
        $item2 = factory(Item::class)->state('published')->make(['user_id' => uuid4()]);
        $item3 = factory(Item::class)->state('draft')->make(['user_id' => $user->id]);
        $item4 = factory(Item::class)->state('draft')->make(['user_id' => uuid4()]);

        $this->assertFalse($user->can('delete', $item1));
        $this->assertFalse($user->can('delete', $item2));
        $this->assertFalse($user->can('delete', $item3));
        $this->assertFalse($user->can('delete', $item4));
    }

    public function test_lolibrarians_can_delete_their_drafts()
    {
        $user = factory(User::class)->state('lolibrarian')->make();

        $item1 = factory(Item::class)->state('draft')->make(['user_id' => $user->id]);

        $this->assertTrue($user->can('delete', $item1));
    }

    public function test_lolibrarians_cannot_delete_other_drafts()
    {
        $user = factory(User::class)->state('lolibrarian')->make();

        $item1 = factory(Item::class)->state('draft')->make(['user_id' => uuid4()]);

        $this->assertFalse($user->can('delete', $item1));
    }

    public function test_lolibrarians_cannot_delete_published_items()
    {
        $user = factory(User::class)->state('lolibrarian')->make();

        $item1 = factory(Item::class)->state('published')->make(['user_id' => $user->id]);
        $item2 = factory(Item::class)->state('published')->make(['user_id' => uuid4()]);

        $this->assertFalse($user->can('delete', $item1));
        $this->assertFalse($user->can('delete', $item2));
    }

    public function test_senior_lolibrarians_can_delete_any_item()
    {
        $user = factory(User::class)->state('senior')->make();

        $item1 = factory(Item::class)->state('published')->make(['user_id' => $user->id]);
        $item2 = factory(Item::class)->state('published')->make(['user_id' => uuid4()]);
        $item3 = factory(Item::class)->state('draft')->make(['user_id' => $user->id]);
        $item4 = factory(Item::class)->state('draft')->make(['user_id' => uuid4()]);

        $this->assertTrue($user->can('delete', $item1));
        $this->assertTrue($user->can('delete', $item2));
        $this->assertTrue($user->can('delete', $item3));
        $this->assertTrue($user->can('delete', $item4));
    }

    public function test_admins_can_delete_any_item()
    {
        $user = factory(User::class)->state('admin')->make();

        $item1 = factory(Item::class)->state('published')->make(['user_id' => $user->id]);
        $item2 = factory(Item::class)->state('published')->make(['user_id' => uuid4()]);
        $item3 = factory(Item::class)->state('draft')->make(['user_id' => $user->id]);
        $item4 = factory(Item::class)->state('draft')->make(['user_id' => uuid4()]);

        $this->assertTrue($user->can('delete', $item1));
        $this->assertTrue($user->can('delete', $item2));
        $this->assertTrue($user->can('delete', $item3));
        $this->assertTrue($user->can('delete', $item4));
    }

    public function test_developers_can_delete_any_item()
    {
        $user = factory(User::class)->state('developer')->make();

        $item1 = factory(Item::class)->state('published')->make(['user_id' => $user->id]);
        $item2 = factory(Item::class)->state('published')->make(['user_id' => uuid4()]);
        $item3 = factory(Item::class)->state('draft')->make(['user_id' => $user->id]);
        $item4 = factory(Item::class)->state('draft')->make(['user_id' => uuid4()]);

        $this->assertTrue($user->can('delete', $item1));
        $this->assertTrue($user->can('delete', $item2));
        $this->assertTrue($user->can('delete', $item3));
        $this->assertTrue($user->can('delete', $item4));
    }

    public function test_lolibrarians_can_delete_items_they_published()
    {
        $user = factory(User::class)->state('lolibrarian')->make();

        $item1 = factory(Item::class)->state('published')->make(['user_id' => $user->id]);
        $item2 = factory(Item::class)->state('published')->make(['user_id' => uuid4()]);

        $this->assertFalse($user->can('delete', $item1));
        $this->assertFalse($user->can('delete', $item2));

        $item1->publisher_id = $user->id;
        $item2->publisher_id = $user->id;

        $this->assertTrue($user->can('delete', $item1));
        $this->assertTrue($user->can('delete', $item2));
    }
}
