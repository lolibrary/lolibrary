<?php

namespace Tests\Feature;

use App\Color;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use Tests\TestCase;

class ColorTest extends TestCase
{
    use DatabaseTransactions;

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testColorIndex()
    {
        $this->seed('ColorSeeder');

        $this->get(route('api.colors.index'))
            ->assertSuccessful()
            ->assertJson(Color::all()->toArray());
    }

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testColorShow()
    {
        $this->seed('ColorSeeder');

        $color = Color::query()->inRandomOrder()->first();

        $this->get(route('api.colors.show', $color))
            ->assertSuccessful()
            ->assertJson($color->toArray());
    }
}
