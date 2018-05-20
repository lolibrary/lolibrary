@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col text-center">
            <img style="max-height: 300px" src="{{ asset('images/banners/banner01.png') }}" alt="">
            <h2>{{ __("Sorry, something on our end broke while loading this page.") }}</h2>
            <h3>{{ __("We've logged the error and will be looking into it!") }}</h3>
        </div>
    </div>
</div>
@endsection
