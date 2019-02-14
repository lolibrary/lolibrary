<?php
namespace App\Http\Controllers;
use App\Http\Controllers\Controller;
use App\Models\Item;
class AdminController extends Controller
{
    /**
     * Get the dashboard view.
     *
     * @return \Illuminate\Http\Response
     */
    public function dashboard()
    {
        $recents = Item::with(Item::PARTIAL_LOAD)
            ->drafts()
            ->orderBy('updated_at', 'desc')
            ->take(10)
            ->get();

        return view('admin.dashboard', compact('recents'));
    }

     /**
     * View all draft items.
     *
     * @return \Illuminate\View\View
     */
    public function queue()
    {
        $items = Item::with(Item::PARTIAL_LOAD)
            ->drafts()
            ->orderBy('updated_at', 'desc')
            ->paginate(30);
        return view('admin.queue', compact('items'));
    }
}
