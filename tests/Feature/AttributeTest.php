<?php

namespace Tests\Feature;

use App\Attribute;

class AttributeTest extends TestCase
{
    /**
     * Test listing all attributes.
     *
     * @return void
     */
    public function testAttributeIndex()
    {
        $this->get(route('api.attributes.index'))
            ->assertSuccessful()
            ->assertJson(Attribute::all()->toArray());
    }

    /**
     * Test listing a specific attribute.
     *
     * @return void
     */
    public function testAttributeShow()
    {
        $attribute = Attribute::query()->inRandomOrder()->first();

        $this->get(route('api.attributes.show', $attribute))
            ->assertSuccessful()
            ->assertJson($attribute->toArray());
    }
}
