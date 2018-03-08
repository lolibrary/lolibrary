<?php

namespace Tests\Feature;

use App\Attribute;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use Tests\TestCase;

class AttributeTest extends TestCase
{
    use DatabaseTransactions;

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testAttributeIndex()
    {
        $this->seed('AttributeSeeder');

        $this->get(route('api.attributes.index'))
            ->assertSuccessful()
            ->assertJson(Attribute::all()->toArray());
    }

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testAttributeShow()
    {
        $this->seed('AttributeSeeder');

        $attribute = Attribute::query()->inRandomOrder()->first();

        $this->get(route('api.attributes.show', $attribute))
            ->assertSuccessful()
            ->assertJson($attribute->toArray());
    }
}
