<?php

namespace Tests\Feature;

use Mockery;
use App\Tag;
use App\Composers\Tags;
use Illuminate\View\View;

class TagComposerTest extends TestCase
{
    /**
     * Test the tags view composer works.
     *
     * @return void
     */
    public function testTagComposing()
    {
        $mock = Mockery::mock(View::class);
        $mock->shouldReceive('with')->once()->andReturnUsing(function ($key, $value) {
            $this->assertEquals('tags', $key);
            $this->assertEquals(Tag::all()->toSelectArray(), $value);
        });

        $composer = new Tags();
        $composer->compose($mock);
    }
}
