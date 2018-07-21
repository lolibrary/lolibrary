<?php

use App\Models\Feature;

class FeatureControllerTest extends TestCase
{
    /**
     * Test listing all features.
     *
     * @return void
     */
    public function testFeatureIndex()
    {
        $this->get(route('api.features.index'))
            ->assertSuccessful()
            ->assertJson(Feature::all()->toArray());
    }

    /**
     * Test getting a specific feature.
     *
     * @return void
     */
    public function testFeatureShow()
    {
        $feature = Feature::query()->inRandomOrder()->first();

        $this->get(route('api.features.show', $feature))
            ->assertSuccessful()
            ->assertJson($feature->toArray());
    }
}
