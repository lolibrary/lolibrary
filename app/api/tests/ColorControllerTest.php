<?php

use App\Models\Color;

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
