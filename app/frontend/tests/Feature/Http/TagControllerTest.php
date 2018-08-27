<?php

namespace Tests\Feature\Http;

use App\Models\Tag;
use Tests\Feature\TestCase;

class TagControllerTest extends TestCase
{
    /**
     * Test the index listing of tags.
     *
     * @return void
     */
    public function testTagIndex()
    {
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
     * Test finding a specific tag.
     *
     * @return void
     */
    public function testTagShow()
    {
        $tag = Tag::query()->inRandomOrder()->first();

        $this->get(route('api.tags.show', $tag))
            ->assertSuccessful()
            ->assertJson($tag->toArray());
    }

    /**
     * Test searching for a tag.
     *
     * @return void
     */
    public function testTagSearch()
    {
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
