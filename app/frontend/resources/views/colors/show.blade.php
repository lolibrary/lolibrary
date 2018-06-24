@extends('layouts.app')

@section('content')
<div class="container">
    @include('components.info', [
        'items' => $items,
        'model' => $color,
        'name' => 'color',
        'image' => false,
        'description' => false,
    ])
    @include('items.listing', ['items' => $items])
</div>
@endsection
