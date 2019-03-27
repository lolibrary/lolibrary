<?php

namespace App\Http\Controllers;

use Auth;
use App\Models\User;
use Illuminate\Http\Request;

class ProfileController extends Controller
{
    /**
     * Construct a new Profile Controller.
     *
     * @return void
     */
    public function __construct()
    {
        $this->middleware('auth');
    }

    /**
     * Show the application dashboard.
     *
     * @return \Illuminate\Http\Response
     */
    public function profile()
    {
        $user = Auth::user();

        return view('profile.index', compact('user'));
    }

    /**
     * Get a user's closet (owned items).
     *
     * @return \Illuminate\Http\Response
     */
    public function closet()
    {
        $user = Auth::user();
        $items = $user->closet()->paginate(24);

        return view('profile.closet', compact('user', 'items'));
    }

    /**
     * Get a user's wishlist (favourited items).
     *
     * @return \Illuminate\Http\Response
     */
    public function wishlist()
    {
        $user = Auth::user();
        $items = $user->wishlist()->paginate(24);

        return view('profile.wishlist', compact('user', 'items'));
    }

    public function edit(String $username)
    {   $profile = User::where('username', $username)->first();
        $user = Auth::user();
        return view('profile.edit', [ 'user' => $profile, 'roles' => User::ROLES]);
    }

    public function destroy(String $username)
    {   $user = User::where('username', $username)->first();
        $user->delete();
        return redirect()
            ->route('admin.users')
            ->with('status', 'User deleted successfully');
    }

    public function save(Request $request, String $username)
    {   $user = User::where('username', $username)->first();
        $user->fill($request->only([
            'email',
            'name',
            'level',
        ]));
        $user->save();
        return redirect()
            ->route('profile.edit', $username)
            ->with('status', 'User saved successfully');
    }
}
