<?php

namespace App\Http\Controllers;

use App\Http\Requests\DonateRequest;

class DonationController extends Controller
{
    /**
     * Redirect to PayPal.
     *
     * @return \Illuminate\Http\RedirectResponse
     */
    public function paypal()
    {
        return redirect(config('services.donation.paypal'));
    }

    /**
     * Redirect to Patreon.
     *
     * @return \Illuminate\Http\RedirectResponse
     */
    public function patreon()
    {
        return redirect(config('services.donation.patreon'));
    }

    /**
     * Get the base donation page.
     *
     * @return \Illuminate\View\View
     */
    public function index()
    {
        return view('donations.donate');
    }

    /**
     * Add a view to say thanks to users for donations.
     *
     * @return \Illuminate\View\View
     */
    public function thanks()
    {
      return view('donations.thanks');
    }
}
