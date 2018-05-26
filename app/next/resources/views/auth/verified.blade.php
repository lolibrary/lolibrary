@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col text-center">
            <img style="max-height: 300px" src="{{ asset('images/banners/banner01.png') }}" alt="">
            <h2>{{ __("Your account is now verified!") }}</h2>
            <h3 class="text-muted">{{ __("You can close this browser tab or go to the homepage.") }}</h3>
        </div>
    </div>
</div>
@endsection
