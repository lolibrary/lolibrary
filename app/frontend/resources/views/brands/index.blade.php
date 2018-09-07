@extends('layouts.app')

@section('content')
<div class="container">
<h2 class="mt-5">{{ __('Brands') }}</h2>
    <div class="flex-grid-container">
        @foreach ($brands as $brand)
            @include( 'components.card', [ 
                'url' => $brand->url, 
                'image' => $brand->image->url, 
                'title' => $brand->name, 
                'size' => '200'
            ])
        @endforeach
</div>

<div class="p-2">
    {{ $brands->links('pagination::simple-bootstrap-4') }}
</div>
@endsection
