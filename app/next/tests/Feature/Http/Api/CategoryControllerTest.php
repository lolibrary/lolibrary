<?php

namespace Tests\Feature\Http\Api;

use App\Models\Category;
use Tests\Feature\TestCase;

class CategoryControllerTest extends TestCase
{
    /**
     * Test listing all categories.
     *
     * @return void
     */
    public function testCategoryIndex()
    {
        $this->get(route('api.categories.index'))
            ->assertSuccessful()
            ->assertJson(Category::all()->toArray());
    }

    /**
     * Test getting a specific category.
     *
     * @return void
     */
    public function testCategoryShow()
    {
        $category = Category::query()->inRandomOrder()->first();

        $this->get(route('api.categories.show', $category))
            ->assertSuccessful()
            ->assertJson($category->toArray());
    }
}
