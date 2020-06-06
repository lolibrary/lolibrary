<?php

namespace App\Http\Controllers\Api;

use App\Models\Attribute;

class AttributeController extends Controller
{
    /**
     * Return all categories, cached.
     *
     * @return mixed
     * @throws \Exception
     */
    public function index()
    {
        return Attribute::cached();
    }

    /**
     * Get a specific category.
     *
     * @param \App\Attribute $attribute
     * @return \App\Attribute
     */
    public function show(Attribute $attribute)
    {
        return $attribute;
    }
}
