@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.hero')

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
    <div style="display: flex; flex-wrap: nowrap; -webkit-overflow-scrolling: touch; overflow-x: auto; overflow-y: hidden">
        @foreach ($brands as $brand)
            <div class="card shadow-sm m-2" style="flex: 0 0 auto">
                <a href="{{ $brand->url }}">
                    <div class="card-body p-0 m-0" style="height: 150px; width: 150px;">
                        <div class="d-flex justify-content-center align-items-center" style="height: 100px">
                            <img src="{{ $brand->image->url }}"
                                onerror="this.src = '{{ asset('images/icon.svg') }}'"
                                class="p-3 mh-100 mw-100" alt=""
                                data-original-url="{{ $brand->image->url }}">
                        </div>
                        <div style="height: 50px" class="text-center">
                            <p class="text-muted small p-2">
                                {{ $brand->name }}
                            </p>
                        </div>
                    </div>
                </a>
            </div>
        @endforeach
    </div>

    <h2 class="mt-5">{{ __('Categories') }}</h2>
    <div style="display: flex; flex-wrap: nowrap; -webkit-overflow-scrolling: touch; overflow-x: auto; overflow-y: hidden">
        @foreach ($categories as $category)
            <div class="card shadow-sm m-2" style="flex: 0 0 auto">
                <a href="{{ $category->url }}">
                    <div class="card-body p-0 m-0" style="height: 150px; width: 150px;">
                        <div class="d-flex justify-content-center align-items-center" style="height: 100px">
                            <img src="{{ asset("categories/{$category->slug}.svg") }}" class="p-1 mh-100 mw-100" alt="">
                        </div>
                        <div style="height: 50px" class="text-center">
                            <p class="text-muted small p-2">
                                {{ $category->name }}
                            </p>
                        </div>
                    </div>
                </a>
            </div>
        @endforeach
    </div>

    <h2 class="mt-5">{{ __('Recent Items') }}</h2>
    <div style="overflow-x: scroll; white-space: nowrap; overflow-y: hidden; -webkit-overflow-scrolling: touch;">
        @foreach ($recent as $item)
            <div class="p-2 d-inline-block" style="width: 16rem; white-space: normal">
                @include('items.card', compact('item'))
            </div>
        @endforeach
    </div>
</div>
@endsection
