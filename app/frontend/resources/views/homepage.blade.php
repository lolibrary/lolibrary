@extends('layouts.app')

@section('content')
<div class="container">
    <h1 class="text-center" style="margin-bottom: 25px">
        <img src="{{ asset('images/banners/banner01.png') }}" alt="" style="max-height: 200px; max-width: 90%">
    </h1>

    <h2 class="sr-only">{{ __('News') }}</h2>
    <div class="row">
        @foreach ($posts as $post)
            <div class="col-md">
                @include('blog.card')
            </div>
        @endforeach
    </div>

    {{-- todo: put brands in here with their images --}}
    {{-- todo: carousel these! (or scroll left/right) --}}
    <h2 class="mt-5">{{ __('Brands') }}</h2>
    <div style="overflow-x: scroll; white-space: nowrap; overflow-y: hidden">
        @foreach ($brands as $brand)
            <div class="p-2 d-inline-block" style="width: 10rem; white-space: normal">
                <div class="card shadow-sm" style="height: 9rem;">
                    <a href="{{ $brand->url }}" class="card-link">
                        <div class="card-body text-center">
                            <div class="mx-auto px-2 py-1">
                                <img crossorigin="anonymous" src="{{ $brand->image->url }}"
                                    onerror="this.src = '{{ asset('images/icon.svg') }}'" class="img-circle mw-100" alt="" style="height: 60px">
                            </div>
                            <p class="text-muted small py-2 align-text-bottom">{{ $brand->name }}</p>
                        </div>
                    </a>
                </div>
            </div>
        @endforeach
    </div>

    <h2 class="mt-5">{{ __('Categories') }}</h2>
    <div style="overflow-x: scroll; white-space: nowrap; overflow-y: hidden">
        @foreach ($categories as $category)
            <div class="p-2 d-inline-block" style="width: 10rem; white-space: normal">
                <div class="card shadow-sm" style="height: 9rem;">
                    <a href="{{ $category->url }}" class="card-link">
                        <div class="card-body text-center">
                            <div class="mx-auto px-2 py-1">
                                <img src="{{ asset('images/icon.svg') }}" class="img-circle mw-100" alt="" style="height: 60px">
                            </div>
                            <p class="text-muted py-2 small align-text-bottom">{{ $category->name }}</p>
                        </div>
                    </a>
                </div>
            </div>
        @endforeach
    </div>

    <h2 class="mt-5">{{ __('Recent Items') }}</h2>
    <div style="overflow-x: scroll; white-space: nowrap; overflow-y: hidden">
        @foreach ($recent as $item)
            <div class="p-2 d-inline-block" style="width: 16rem; white-space: normal">
                @include('items.card', compact('item'))
            </div>
        @endforeach
    </div>
</div>
@endsection
