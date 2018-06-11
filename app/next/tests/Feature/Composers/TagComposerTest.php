<?php

namespace Tests\Feature\Composers;

use Mockery;
use App\Models\Tag;
use App\Composers\Tags;
use Illuminate\View\View;
use Tests\Feature\TestCase;

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
