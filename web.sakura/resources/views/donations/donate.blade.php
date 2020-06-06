@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.hero')

    <div class="row p-0">
        <div class="col-sm-10 col-md-8 col-lg-6 offset-sm-1 offset-md-2 offset-lg-3 px-2 text-center">
          <p>Lolibrary is funded entirely by donations from our users,
            and we're eternally grateful for all the support you give us!</p>
          <p>We're a registered non-profit, and all funds will go towards operating costs for Lolibrary, as well as future development.
            If you'd prefer to <a href="https://patreon.com/lolibrary" target="_blank" rel="external">support us on Patreon</a>, you can go there, too!</p>

          <p class="my-4">
            <a title="Support us on patreon" href="https://patreon.com/lolibrary" target="_blank" rel="external">
              <img src="{{ cdn_link('assets/become_a_patron_button@2x.png') }}"
                style="max-width: 217px; max-height: 51px">
            </a>
          </p>

          <p>Alternatively, you can donate using the link below, where you can pay with Card, PayPal or Apple Pay.</p>

          <p class="my-4">
            <a href="https://donorbox.org/lolibrary" target="_blank" rel="external" class="mx-3 btn btn-lg btn-outline-primary">
              Donate
            </a>
          </p>

          <p class="small text-muted">Lolibrary is a 501(c)(3) registered non-profit incorporated in the USA; all donations are tax-deductible and our EIN is 81-2942481.</p>

        </div>
    </div>

</div>
@endsection

@section('meta')
    <link rel="canonical" href="{{ route('donate') }}">

    <meta property="og:url" content="{{ route('donate') }}">
    <meta property="og:type" content="article">
    <meta property="og:title" content="Donate to Lolibrary">
    <meta property="og:image" content="{{ cdn_link('assets/banners/banner01-white.png') }}">
@endsection
