@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.hero')

    {{-- todo: put brands in here with their images --}}
    {{-- todo: carousel these! (or scroll left/right) --}}
    <h2 class="mt-5">{{ __('Brands') }}</h2>
    <div class="scrollbox">
        @foreach ($brands as $brand)
        <div class="scrollbox-item m-2">
            <div class="card shadow-sm scrollbox-square">
                <a href="{{ $brand->url }}">
                    <div class="scrollbox-img">
                        <img src="{{ Storage::cloud()->url($brand->image) }}" alt="" data-original-url="{{ Storage::cloud()->url($brand->image) }}"
                            onerror="if (this.src !== '{{ cdn_link('categories/other.svg') }}') this.src = '{{ cdn_link('categories/other.svg') }}'">
                    </div>
                    <div class="scrollbox-text">
                        <p class="text-muted small p-0 m-0">{{ $brand->name }}</p>
                    </div>
                </a>
            </div>
        </div>
        @endforeach
    </div>

    <h2 class="mt-5">{{ __('Categories') }}</h2>
    <div class="scrollbox">
        @foreach ($categories as $category)
        <div class="scrollbox-item m-2">
            <div class="card shadow-sm scrollbox-square">
                <a href="{{ $category->url }}">
                    <div class="scrollbox-img">
                        <img src="{{ Storage::cloud()->url($category->image) }}" alt="">
                    </div>
                    <div class="scrollbox-text">
                        <p class="text-muted small p-0 m-0">{{ $category->name }}</p>
                    </div>
                </a>
            </div>
        </div>
        @endforeach
    </div>

    <h2 class="mt-5">{{ __('Recent Items') }}</h2>
    <div class="scrollbox">
        @foreach ($recent as $item)
            <div class="scrollbox-item scrollbox-item-card m-2">
                @include('items.card', compact('item'))
            </div>
        @endforeach
    </div>
</div>
@endsection
