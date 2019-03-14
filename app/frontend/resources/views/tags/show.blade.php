@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.info', [
        'items' => $items,
        'model' => $tag,
        'name' => 'tag',
        'tag' => true,
    ])
    @include('items.listing', ['items' => $items])
</div>
@endsection

@section('meta')
    <link rel="canonical" href="{{ $tag->url }}">

    <meta property="og:url" content="{{ $tag->url }}">
    <meta property="og:type" content="article">
    <meta property="og:title" content="Tag '{{ $tag->name }}' on Lolibrary">
    <meta property="og:image" content="{{ cdn_link('assets/banners/banner01-white.png') }}">
@endsection
