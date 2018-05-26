<?php

namespace Tests\Feature;

use Mockery;
use App\Feature;
use App\Composers\Features;
use Illuminate\View\View;

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
