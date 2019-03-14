@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.hero')

    <div class="row p-0">
        <div class="col-sm-10 col-md-8 col-lg-6 offset-sm-1 offset-md-2 offset-lg-3 px-2 text-center">
          <h1 class="h2">Thank you for donating to Lolibrary!</h1>

          <p class="lead" >It's folks like you who enable us to keep offering the service we do. Thank you!</p>

          <p class="mt-4 lead">Donate regularly?</p>
          <p class="my-4">
            <a title="Support us on patreon" href="{{ config('services.donation.patreon') }}" target="_blank" rel="external">
              <img src="{{ cdn_link('assets/become_a_patron_button@2x.png') }}"
                style="max-width: 217px; max-height: 51px">
            </a>
          </p>
          <p class="text-muted">Patrons on patreon can get extra goodies for donating monthly!</p>
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
