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
        return redirect()->route('search');
    }

    /**
     * Update a user's wishlist.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\RedirectResponse
     */
    public function wishlist(Item $item)
    {
        $user = auth()->user();
        $attached = $user->updateWishlist($item);
        $status = $attached ? 'added' : 'removed';

        return back()->withStatus(trans("user.wishlist.{$status}", ['item' => str_limit($item->english_name, 28)]));
    }

    /**
     * Update a user's closet.
     *
     * @param \App\Models\Item $item
     * @return \Illuminate\Http\RedirectResponse
     */
    public function closet(Item $item)
    {
        $user = auth()->user();
        $attached = $user->updateCloset($item);
        $status = $attached ? 'added' : 'removed';

        return back()->withStatus(trans("user.closet.{$status}", ['item' => str_limit($item->english_name, 28)]));
    }
}
