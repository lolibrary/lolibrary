<?php

namespace Tests\Feature;

use App\Feature;
use Tests\TestCase;
use Illuminate\Foundation\Testing\DatabaseMigrations;

class FeatureTest extends TestCase
{
    use DatabaseMigrations;

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testFeatureIndex()
    {
        $features = factory(Feature::class, 5)->create();

        $this->get(route('api.features.index'))
            ->assertSuccessful()
            ->assertJson($features->toArray());
    }
}
