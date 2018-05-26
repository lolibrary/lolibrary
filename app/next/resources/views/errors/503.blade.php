@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col text-center">
            <img style="max-height: 300px" src="{{ asset('images/banners/banner01.png') }}" alt="">
            <h2>{{ __("We are doing a little spring cleaning.") }}</h2>
            <h3 class="text-muted">{{ __("Be right back!") }}</h3>
        </div>
    </div>
</div>
@endsection
