@extends('layouts.app')

@section('content')
<div class="container">
<h2 class="mt-5">{{ __('Categories') }}</h2>
    <div class="flex-grid-container">
        @foreach ($categories as $category)
            @include( 'components.card', [ 
                'url' => $category->url, 
                'image' => asset("categories/{$category->slug}.svg"), 
                'title' => $category->name, 
                'size' => '200'
            ])
        @endforeach
</div>

<div class="p-2">
    {{ $categories->links('pagination::simple-bootstrap-4') }}
</div>
@endsection
