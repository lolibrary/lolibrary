<!DOCTYPE html>
<html lang="{{ app()->getLocale() }}">
<head prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# product: http://ogp.me/ns/product#">
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="icon" href="{{ asset('images/favicon.png') }}" type="image/png" sizes="any">
    <title>{{ $title or config('app.name', 'Lolibrary') }}</title>

    <!-- CSRF Token -->
    <meta name="csrf-token" content="{{ csrf_token() }}">
    <meta name="search-endpoint" content="{{ route('api.search') }}">

    @include('components.external.font-awesome')

    <!-- Styles -->
    <link href="{{ asset('assets/app.css') }}" rel="stylesheet">
</head>
<body>
    <a class="sr-only sr-only-focusable" href="#skip-navigation">{{ __('Skip to content') }}</a>
    <div id="app" style="margin-top: 55px">
        <nav class="navbar navbar-expand-md navbar-light navbar-laravel fixed-top">
            <div class="container">
                <a class="navbar-brand" href="{{ url('/') }}">
                    <img style="height: 14px" src="{{ asset('images/logo_horizontal.png') }}" alt="Lolibrary logo">
                </a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>

                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <!-- Left Side Of Navbar -->
                    <ul class="navbar-nav mr-auto">

                    </ul>

                    <!-- Right Side Of Navbar -->
                    <ul class="navbar-nav ml-auto">
                        <!-- Authentication Links -->
                        @guest
                            <li><a class="nav-link" href="{{ route('login') }}">{{ __('Login') }}</a></li>
                            <li><a class="nav-link" href="{{ route('register') }}">{{ __('Register') }}</a></li>
                        @else
                            @include('components.navbar.dropdown')
                        @endguest
                    </ul>
                </div>
            </div>
        </nav>

        @if (config('app.banner.show'))
            <div class="alert-fullwidth text-center alert alert-{{ config('app.banner.style', 'info') }}" role="alert">
                {{ config('app.banner.content') }}
            </div>
        @endif

        @if (session('status'))
            <div class="alert-fullwidth text-center alert alert-primary" role="alert">
                {{ session('status') }}
            </div>
        @endif

        <main class="py-4" id="skip-navigation">
            @yield('content')
        </main>
    </div>

    <!-- Scripts -->
    <script src="{{ asset('assets/app.js') }}"></script>
</body>
</html>
