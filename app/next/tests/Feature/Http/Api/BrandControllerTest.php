<?php

namespace Tests\Feature\Http\Api;

use App\Brand;
use Tests\Feature\TestCase;

class BrandControllerTest extends TestCase
{
    /**
     * Test listing all brands.
     *
     * @return void
     */
    public function testBrandIndex()
    {
        $this->get(route('api.brands.index'))
            ->assertSuccessful()
            ->assertJson(Brand::all()->toArray());
    }

    /**
     * Test listing a specific brand.
     *
     * @return void
     */
    public function testBrandShow()
    {
        $brand = Brand::query()->inRandomOrder()->first();

        $this->get(route('api.brands.show', $brand))
            ->assertSuccessful()
            ->assertJson($brand->toArray());
    }
}
