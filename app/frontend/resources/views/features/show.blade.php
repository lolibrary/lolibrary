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
