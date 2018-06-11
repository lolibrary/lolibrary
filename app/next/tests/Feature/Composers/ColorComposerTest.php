<?php

namespace Tests\Feature\Composers;

use Mockery;
use App\Models\Color;
use App\Composers\Colors;
use Illuminate\View\View;
use Tests\Feature\TestCase;

class ColorComposerTest extends TestCase
{
    /**
     * Test the colors view composer works.
     *
     * @return void
     */
    public function testColorComposing()
    {
        $mock = Mockery::mock(View::class);
        $mock->shouldReceive('with')->once()->andReturnUsing(function ($key, $value) {
            $this->assertEquals('colors', $key);
            $this->assertEquals(Color::all()->toSelectArray(), $value);
        });

        $composer = new Colors();
        $composer->compose($mock);
    }
}
