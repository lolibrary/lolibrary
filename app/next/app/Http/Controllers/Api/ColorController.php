<?php

namespace App\Http\Controllers\Api;

use App\Models\Color;

class ColorController extends Controller
{
    /**
     * Return all categories, cached.
     *
     * @return mixed
     * @throws \Exception
     */
    public function index()
    {
        return Color::cached();
    }

    /**
     * Get a specific category.
     *
     * @param \App\Color $color
     * @return \App\Color
     */
    public function show(Color $color)
    {
        return $color;
    }
}
