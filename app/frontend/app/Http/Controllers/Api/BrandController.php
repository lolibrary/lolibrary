<?php

namespace App\Http\Controllers\Api;

use App\Models\Brand;

class BrandController extends Controller
{
    /**
     * Return all brands, cached.
     *
     * @return mixed
     * @throws \Exception
     */
    public function index()
    {
        return Brand::cached();
    }

    /**
     * Get a specific category.
     *
     * @param \App\Brand $brand
     * @return \App\Brand
     */
    public function show(Brand $brand)
    {
        return $brand;
    }
}
