<?php

namespace Tests\Feature;

use App\Category;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use Tests\TestCase;

class CategoryTest extends TestCase
{
    use DatabaseTransactions;

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testCategoryIndex()
    {
        $this->seed('CategorySeeder');

        $this->get(route('api.categories.index'))
            ->assertSuccessful()
            ->assertJson(Category::all()->toArray());
    }

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testCategoryShow()
    {
        $this->seed('CategorySeeder');

        $category = Category::query()->inRandomOrder()->first();

        $this->get(route('api.categories.show', $category))
            ->assertSuccessful()
            ->assertJson($category->toArray());
    }
}
