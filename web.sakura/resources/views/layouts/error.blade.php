@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col text-center">
            <img style="max-height: 300px; max-width: 100%" src="{{ cdn_link('assets/banners/banner01.png') }}" alt="">
            @yield('error')
        </div>
    </div>
</div>
@endsection
