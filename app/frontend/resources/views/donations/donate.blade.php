@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.hero')

    <div class="lead mb-4 text-center">
        <p>Lolibrary is a 501(c)(3) non-profit run off of your donations.</p>
        <p>We are dedicated to archiving and preserving alternative street fashion.</p>
    </div>

    <div class="row mb-4">
        <div class="col-sm-6 col-md-5 col-lg-4 col-xl-3 offset-xl-3 offset-lg-2 offset-md-1 offset-sm-0">
            <a href="{{ route('donate.paypal') }}" class="btn btn-block btn-default btn-lg text-white mb-3" style="background-color: #0070ba"><i class="fab fa-paypal"></i> Donate with PayPal</a>
        </div>
        <div class="col-sm-6 col-md-5 col-lg-4 col-xl-3">
            <a href="{{ route('donate.patreon') }}" class="btn btn-block btn-default btn-lg text-white" style="background-color: rgb(249, 104, 84)"><i class="fab fa-patreon"></i> Support us on Patreon</a>
        </div>
    </div>
    <div class="row">
        <div class="col-sm-12 col-md-10 col-lg-8 col-xl-6 mx-auto text-center">
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
        <form id="donationForm" method="post">
            @csrf
            <input type="hidden" id="stripeToken" name="stripeToken">
            <input type="hidden" name="amount" id="amount">
            <div class="row text-center mx-auto">
                <div class="col-md-9 mx-auto">
                    <div class="row">
                        <div class="col-4 px-2">
                            <button disabled class="btn btn-lg btn-block btn-outline-dark">US $</i></button>
                        </div>
                        <div class="col-8 pr-2 pl-0" id="donationFeedbackContainer">
                            <input class="form-control form-control-lg" value="1.00" placeholder="1.00 USD" id="donationAmount" autocomplete="off">
                        </div>
                    </div>
                </div>
            </div>
            <p class="text-danger text-center mt-2" id="donationFeedback" style="display: none">
                Should be a value above <span class="text-regular">US$1.00</span>
            </p>
            <p class="text-center mt-3">
                <button type="submit" class="btn btn-outline-primary btn-lg" id="donate">Donate</button>
            </p>
        </form>
      </div>
    </div>
  </div>
</div>
@endsection

@section('script')
<script src="https://checkout.stripe.com/checkout.js"></script>
<script>

var handler = StripeCheckout.configure({
    key: "{{ config('services.stripe.key') }}",
    locale: 'auto',
    name: 'Lolibrary Inc',
    description: 'One-time donation',
    image: '/images/icon.png',
    token: function(token) {
        $('input#stripeToken').val(token.id);
        $('form#donationForm').submit();
    }
});

var donationForm = document.getElementById('donationForm');

donationForm.addEventListener('submit', function (event) {
    event.preventDefault();
    event.stopPropagation();

    var valid = validateCurrencyForm();
    var amount = parseFloat($('input#donationAmount').val());

    if (valid === true) {
        amount = Math.round(amount * 100);
        
        $('input#amount').val(amount);

        handler.open({
            amount: amount
        });
    }
});

// temporary until it's in app.js
function validateCurrencyForm() {
    var $input = $('input#donationAmount');
    var $container = $('#donationFeedbackContainer');
    var $feedback = $('#donationFeedback');

    if (validateCurrency($input.val()) === false) {
        $input.removeClass('is-valid').addClass('is-invalid');
        $feedback.show();

        return false;
    }

    $input.removeClass('is-invalid').addClass('is-valid');
    $feedback.hide();

    return true;
}

function validateCurrency(amount) {
    var parsed = parseFloat(amount);

    return !isNaN(parsed) &&
        parsed >= 1.00 &&
        parsed <= 9999 &&
        parsed == amount;
}
</script>
@endsection

@section('meta')
    <link rel="canonical" href="{{ route('donate') }}">

    <meta property="og:url" content="{{ route('donate') }}">
    <meta property="og:type" content="article">
    <meta property="og:title" content="Donate to Lolibrary">
    <meta property="og:image" content="{{ asset('images/banners/banner01-white.png') }}">
@endsection
