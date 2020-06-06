@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.info', [
        'items' => $items,
        'model' => $category,
        'name' => 'category',
        'image' => asset("categories/{$category->slug}.svg"),
        'description' => true,
    ])
    @include('items.listing', ['items' => $items])
</div>
@endsection

@section('meta')
    <link rel="canonical" href="{{ $category->url }}">

    <meta property="og:url" content="{{ $category->url }}">
    <meta property="og:type" content="article">
    <meta property="og:title" content="Lolibrary: {{ $category->name }}">
    <meta property="og:image" content="{{ $category->image->url ?? default_asset() }}">
@endsection
