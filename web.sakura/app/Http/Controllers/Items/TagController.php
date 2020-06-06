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
        return redirect()->to(search_route(['tags' => [$tag->slug]]));
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
