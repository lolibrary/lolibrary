@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col text-center">
            <img style="max-height: 300px" src="{{ asset('images/banners/banner01.png') }}" alt="">
            <h2>{{ __("Sorry, the page you are looking for could not be found.") }}</h2>
        </div>
    </div>
</div>
@endsection
