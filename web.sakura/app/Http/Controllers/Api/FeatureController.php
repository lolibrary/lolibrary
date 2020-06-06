<?php

namespace App\Http\Controllers\Api;

use App\Models\Feature;

class FeatureController extends Controller
{
    /**
     * Return all categories, cached.
     *
     * @return mixed
     * @throws \Exception
     */
    public function index()
    {
        return Feature::cached();
    }

    /**
     * Get a specific category.
     *
     * @param \App\Feature $feature
     * @return \App\Feature
     */
    public function show(Feature $feature)
    {
        return $feature;
    }
}
