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
        $items = $category->items()->with(Item::PARTIAL_LOAD)->paginate(24);

        return view('categories.show', compact('category', 'items'));
    }

    /**
     * Show a paginated list of categories.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // todo: make this a static ::index() method.
        $categories = Category::paginate(36);

        return view('categories.index', compact('categories'));
    }
}
