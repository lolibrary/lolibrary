<?php

namespace App\Http\Controllers;

use Auth;
use App\Models\User;
use App\Notifications\VerifyEmail;

class EmailController extends Controller
{
    /**
     * Construct a new email controller for verification.
     */
    public function __construct()
    {
        $this->middleware('auth')->only('resend');
    }

    /**
     * Verify a user's email address.
     *
     * @return \Illuminate\Http\Response
     */
    public function verify(string $email, string $token)
    {
        $user = User::email($email)->firstOrFail();

        if ($user->verified === false && hash_equals($user->email_token, $token)) {
            $user->email_token = null;
            $user->save();

            return view('auth.verified');
        }

        abort(404);
    }

    /**
     * Verify a user's email address.
     *
     * @return \Illuminate\Http\Response
     */
    public function resend()
    {
        $user = Auth::user();

        if ($user->verified) {
            return back();
        }

        $user->notify(new VerifyEmail($user));
    }
}
