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
     * Redirect to the search page.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return redirect()->route('search');
    }
}
