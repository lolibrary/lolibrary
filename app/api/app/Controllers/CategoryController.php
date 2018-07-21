<?php

namespace App\Controllers;

use App\Models\Category;

class CategoryController extends Controller
{
    /**
     * Return all categories, cached.
     *
     * @return mixed
     * @throws \Exception
     */
    public function index()
    {
        return Category::cached();
    }

    /**
     * Get a specific category.
     *
     * @param \App\Category $category
     * @return \App\Category
     */
    public function show(Category $category)
    {
        return $category;
    }
}
