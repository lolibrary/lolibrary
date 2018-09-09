<?php

namespace App\Http\Controllers\Items;

use App\Models\Item;
use App\Models\Category;
use App\Http\Controllers\Controller;

class CategoryController extends Controller
{
    /**
     * Show a category.
     *
     * @param \App\Models\Category $category
     * @return \Illuminate\Http\Response
     */
    public function show(Category $category)
    {
        return redirect()->to(search_route(['categories' => [$category->slug]]));
    }

    /**
     * Redirect to the search page.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return redirect()->route('search');
    }
}
