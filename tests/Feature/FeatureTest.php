<?php

namespace Tests\Feature;

use App\Feature;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use Tests\TestCase;

class FeatureTest extends TestCase
{
    use DatabaseTransactions;

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testFeatureIndex()
    {
        $this->seed('FeatureSeeder');

        $this->get(route('api.features.index'))
            ->assertSuccessful()
            ->assertJson(Feature::all()->toArray());
    }

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testFeatureShow()
    {
        $this->seed('FeatureSeeder');

        $feature = Feature::query()->inRandomOrder()->first();

        $this->get(route('api.features.show', $feature))
            ->assertSuccessful()
            ->assertJson($feature->toArray());
    }
}
