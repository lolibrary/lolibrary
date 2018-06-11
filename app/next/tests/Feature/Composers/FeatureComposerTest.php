<?php

namespace Tests\Feature\Composers;

use Mockery;
use App\Models\Feature;
use Illuminate\View\View;
use App\Composers\Features;
use Tests\Feature\TestCase;

class FeatureComposerTest extends TestCase
{
    /**
     * Test the features view composer works.
     *
     * @return void
     */
    public function testFeatureComposing()
    {
        $mock = Mockery::mock(View::class);
        $mock->shouldReceive('with')->once()->andReturnUsing(function ($key, $value) {
            $this->assertEquals('features', $key);
            $this->assertEquals(Feature::all()->toSelectArray(), $value);
        });

        $composer = new Features();
        $composer->compose($mock);
    }
}
