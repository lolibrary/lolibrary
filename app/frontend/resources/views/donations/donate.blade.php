@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.hero')

    <div class="row mb-4">
        <div class="col-sm-4 offset-sm-2 mb-4">
            <a href="{{ route('donate.paypal') }}" class="btn btn-block btn-default btn-lg text-white" style="background-color: #0070ba"><i class="fab fa-paypal"></i> Donate with PayPal</a>
        </div>
        <div class="col-sm-4">
            <a href="{{ route('donate.patreon') }}" class="btn btn-block btn-default btn-lg text-white" style="background-color: rgb(249, 104, 84)"><i class="fab fa-patreon"></i> Support us on Patreon</a>
        </div>
    </div>
    <div class="row">
        <div class="col-sm-5 mx-auto text-center">
            <div class="card">
                <div class="card-body">
                    <p>
                        You can also choose to pay with Apple Pay, Google Pay,<br> or Credit/Debit Card.
                    </p>
                    <p>
                        Just click below and enter your chosen donation amount.
                    </p>
                    <p class="py-2">
                        <a href="#" class="btn btn-lg btn-outline-primary" data-toggle="modal" data-target="#donationEntry">Choose Amount</a>
                    </p>
                    <a href="https://stripe.com" rel="external nofollow">
                        <img alt="Powered by Stripe" src="{{ asset('images/powered_by_stripe.svg') }}">
                    </a>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="modal" tabindex="-1" role="dialog" id="donationEntry">
  <div class="modal-dialog modal-dialog-centered" role="document">
    <div class="modal-content">
      <div class="modal-body">
        <div class="mb-4">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
        </div>
        <p class="lead text-center">Thank you for donating to Lolibrary!</p>
        <p class="text-center px-3 text-muted">Simply enter how much you'd like to donate below, and you'll be asked for details 💖</p>
        <div class="row p-4">
            <div class="col-3 px-2 text-right">
                <button disabled class="btn btn-outline-dark"><i class="far fa-dollar-sign align-text-bottom" style="font-size: 1.75rem; padding-top: .25rem"></i></button>
            </div>
            <div class="col-8 px-2 text-left">
                <input class="form-control form-control-lg" value="1.00" placeholder="1.00 USD">
            </div>
        </div>
        <p class="text-center">
            <button class="btn btn-outline-primary btn-lg">Donate</button>
        </p>
        <p class="text-center small text-muted">The minimum you can pay this way is <span class="text-regular">US$1.00</span> due to fees</p>
      </div>
    </div>
  </div>
</div>
@endsection
