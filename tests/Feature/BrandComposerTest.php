<?php

namespace Tests\Feature;

use Mockery;
use App\Brand;
use App\Composers\Brands;
use Illuminate\View\View;

class BrandComposerTest extends TestCase
{
    /**
     * Test the brand view composer works.
     *
     * @return void
     */
    public function testBrandComposing()
    {
        $mock = Mockery::mock(View::class);
        $mock->shouldReceive('with')->once()->andReturnUsing(function ($key, $value) {
            $this->assertEquals('brands', $key);
            $this->assertEquals(Brand::all()->toSelectArray('short_name'), $value);
        });

        $composer = new Brands();
        $composer->compose($mock);
    }
}
