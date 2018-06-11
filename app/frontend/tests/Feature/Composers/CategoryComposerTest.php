<?php

namespace Tests\Feature\Composers;

use Mockery;
use App\Models\Category;
use Illuminate\View\View;
use Tests\Feature\TestCase;
use App\Composers\Categories;

class CategoryComposerTest extends TestCase
{
    /**
     * Test the categories view composer works.
     *
     * @return void
     */
    public function testCategoryComposing()
    {
        $mock = Mockery::mock(View::class);
        $mock->shouldReceive('with')->once()->andReturnUsing(function ($key, $value) {
            $this->assertEquals('categories', $key);
            $this->assertEquals(Category::all()->toSelectArray(), $value);
        });

        $composer = new Categories();
        $composer->compose($mock);
    }
}
