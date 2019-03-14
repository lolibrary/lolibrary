@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.info', [
        'items' => $items,
        'model' => $feature,
        'name' => 'feature',
        'description' => true,
    ])
    @include('items.listing', ['items' => $items])
</div>
@endsection

@section('meta')
    <link rel="canonical" href="{{ $feature->url }}">

    <meta property="og:url" content="{{ $feature->url }}">
    <meta property="og:type" content="article">
    <meta property="og:title" content="Feature '{{ $feature->name }}' on Lolibrary">
    <meta property="og:image" content="{{ cdn_link('assets/banners/banner01-white.png') }}">
@endsection
