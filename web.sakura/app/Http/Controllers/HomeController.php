<?php

namespace App\Http\Controllers;

use App\Models\Post;
use App\Models\Item;
use App\Models\Brand;
use App\Models\Category;

class HomeController extends Controller
{
    /**
     * Show the application dashboard.
     *
     * @return \Illuminate\Http\Response
     */
    public function homepage()
    {
        // todo: make this a static ::homepage() method
        $posts = Post::query()
            ->with('user')
            ->whereNotNull('published_at')
            ->take(3)
            ->orderBy('published_at', 'desc')
            ->get();

        $brands = Brand::all();
        $categories = Category::all();
        $recent = Item::with(Item::PARTIAL_LOAD)
            ->drafts(false)
            ->orderBy('published_at', 'desc')
            ->take(15)
            ->get();

        return view('homepage', compact('posts', 'brands', 'categories', 'recent'));
    }
}
