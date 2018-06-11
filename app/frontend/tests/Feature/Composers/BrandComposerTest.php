<?php

namespace Tests\Feature\Composers;

use Mockery;
use App\Models\Brand;
use App\Composers\Brands;
use Illuminate\View\View;
use Tests\Feature\TestCase;

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
            $this->assertEquals(Brand::all()->toSelectArray(), $value);
        });

        $composer = new Brands();
        $composer->compose($mock);
    }
}
