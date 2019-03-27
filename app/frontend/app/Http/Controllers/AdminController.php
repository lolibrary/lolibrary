<?php
namespace App\Http\Controllers;
use App\Http\Controllers\Controller;
use App\Models\Item;
use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Database\Eloquent\Builder;

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

    public function users(Request $request)
    {
        $users = User::query();

        if (is_string($request->search) && strlen($request->search) > 0){
            $search = '%' . $request->search . '%';

            $users = $users->where(function (Builder $users) use ($search) {
                $users->where('name', 'ilike', $search);
                $users->orWhere('username', 'ilike', $search);
                $users->orWhere('email', 'ilike', $search);
            });
        }

        $users = $users->orderBy('username', 'asc')->paginate(30);
        return view('admin.users', compact('users'));
    }
}
