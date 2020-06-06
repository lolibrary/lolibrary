@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.info', [
        'items' => $items,
        'model' => $brand,
        'name' => 'brand',
        'image' => true,
        'description' => true,
    ])
    @include('items.listing', ['items' => $items])
</div>
@endsection
