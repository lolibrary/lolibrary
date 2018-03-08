<?php

namespace Tests\Feature;

use App\Brand;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use Tests\TestCase;

class BrandTest extends TestCase
{
    use DatabaseTransactions;

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testBrandIndex()
    {
        $this->seed('BrandSeeder');

        $this->get(route('api.brands.index'))
            ->assertSuccessful()
            ->assertJson(Brand::all()->toArray());
    }

    /**
     * A basic test example.
     *
     * @return void
     */
    public function testBrandShow()
    {
        $this->seed('BrandSeeder');

        $brand = Brand::query()->inRandomOrder()->first();

        $this->get(route('api.brands.show', $brand))
            ->assertSuccessful()
            ->assertJson($brand->toArray());
    }
}
