<?php

namespace App\Http\Controllers\Items;

use App\Models\Item;
use App\Http\Controllers\Controller;

class ItemController extends Controller
{
    /**
     * Show an item.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\Response
     */
    public function show(Item $item)
    {
        $item->load(Item::FULLY_LOAD);

        return view('items.show', compact('item'));
    }

    /**
     * Show a paginated list of items.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // todo: make this a static ::index() method.
        $items = Item::paginate(52);

        return view('items.index', compact('items'));
    }
}
