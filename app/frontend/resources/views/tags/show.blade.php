@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row text-center">
        <div class="p-1 list-group text-center col-lg-4 col-sm-6 small p-2">
            <a class="list-group-item" href="{{ $tag->url }}">
                #{{ $tag->slug }}
            </a>
        </div>
        <p class="col-lg-8 col-sm-6 mt-3">
            <span class="text-regular">{{ $items->total() }}</span> items found with this tag.
        </p>
    </div>

    <div class="row">
        @forelse ($items as $item)
            <div class="col-lg-3 col-md-4 col-sm-6 p-2">
                @include('items.card', ['item' => $item])
            </div>
        @empty
            <p class="text-center">No items found!</p>
        @endif
    </div>

    <div class="p-2">
        {{ $items->links('pagination::simple-bootstrap-4') }}
    </div>
</div>
@endsection
