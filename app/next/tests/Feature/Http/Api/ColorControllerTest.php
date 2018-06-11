<?php

namespace Tests\Feature\Http\Api;

use App\Models\Color;
use Tests\Feature\TestCase;

class ColorControllerTest extends TestCase
{
    /**
     * Test listing all colorways.
     *
     * @return void
     */
    public function testColorIndex()
    {
        $this->get(route('api.colors.index'))
            ->assertSuccessful()
            ->assertJson(Color::all()->toArray());
    }

    /**
     * Test listing a specific colorway.
     *
     * @return void
     */
    public function testColorShow()
    {
        $color = Color::query()->inRandomOrder()->first();

        $this->get(route('api.colors.show', $color))
            ->assertSuccessful()
            ->assertJson($color->toArray());
    }
}
