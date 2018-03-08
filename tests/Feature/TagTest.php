<?php

namespace Tests\Feature;

use App\Tag;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use Tests\TestCase;

class TagTest extends TestCase
{
    use DatabaseTransactions;

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testTagIndex()
    {
        $this->seed('TagSeeder');

        $this->get(route('api.tags.index'))
            ->assertSuccessful()
            ->assertJsonStructure([
                'current_page',
                'data' => ['*' => ['name', 'slug']],
                'first_page_url',
                'from',
                'last_page',
                'last_page_url',
                'next_page_url',
                'path',
                'per_page',
                'prev_page_url',
                'to',
                'total',
            ]);
    }

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testTagShow()
    {
        $this->seed('TagSeeder');

        $tag = Tag::query()->inRandomOrder()->first();

        $this->get(route('api.tags.show', $tag))
            ->assertSuccessful()
            ->assertJson($tag->toArray());
    }

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testTagSearch()
    {
        $this->seed('TagSeeder');

        $this->get(route('api.tags.search') . '?search=tag')
            ->assertSuccessful()
            ->assertJsonStructure([
                'current_page',
                'data' => ['*' => ['name', 'slug']],
                'first_page_url',
                'from',
                'last_page',
                'last_page_url',
                'next_page_url',
                'path',
                'per_page',
                'prev_page_url',
                'to',
                'total',
            ]);
    }
}
