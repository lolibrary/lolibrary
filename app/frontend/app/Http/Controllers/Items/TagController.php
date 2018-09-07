<?php

namespace App\Http\Controllers\Items;

use App\Models\Tag;
use App\Models\Item;
use App\Http\Controllers\Controller;

class TagController extends Controller
{
    /**
     * Show a tag.
     *
     * @param \App\Models\Tag $tag
     * @return \Illuminate\Http\Response
     */
    public function show(Tag $tag)
    {
        $items = $tag->items()->with(Item::PARTIAL_LOAD)->paginate(24);

        return view('tags.show', compact('tag', 'items'));
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
