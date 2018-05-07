<?php

namespace Tests\Unit;

use Mockery;
use Tests\TestCase;
use App\Composers\Years;
use Illuminate\View\View;

class YearComposerTest extends TestCase
{
    /**
     * Test the years view composer works properly.
     *
     * @return void
     */
    public function testYearsComposer()
    {
        $mock = Mockery::mock(View::class);
        $mock->shouldReceive('with')->once()->andReturnUsing(function ($key, $value) {
            $this->assertEquals('years', $key);
            $this->assertEquals(array_reverse(range(1990, date('Y') + 3)), $value);
        });

        $composer = new Years();
        $composer->compose($mock);
    }
}
