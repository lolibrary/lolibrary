@extends('profile.layout')

@section('profile')
@if ($items->count() > 0)
    <div class="row">
        @foreach ($items as $item)
            @component('items.card', ['item' => $item])
                <button class="btn btn-outline-danger btn-block rounded-0" style="border: none; border-top: 1px solid rgba(0, 0, 0, 0.125);" href="{{ $item->url }}">
                    Remove from Wishlist
                </button>
            @endcomponent
        @endforeach
    </div>

    {{ $items->links() }}
@else
<div class="text-center mt-5">
    <p class="h2">There are no items in your wishlist.</p>
    <p class="lead">Why not <a href="{{ route('search') }}">search for some items to add</a>?</p>
</div>
<div class="row pt-5">
    <div class="col-4">
        <div class="card bg-light text-muted">
            <div class="card-body shadow-sm d-flex justify-content-center align-items-center">
                <i class="fal fa-plus-circle fa-5x"></i>
            </div>
        </div>
    </div>
    <div class="col-4">
        <div class="card bg-light text-muted">
            <div class="card-body shadow-sm d-flex justify-content-center align-items-center">
                <i class="fal fa-plus-circle fa-5x"></i>
            </div>
        </div>
    </div>
    <div class="col-4">
        <div class="card bg-light text-muted">
            <div class="card-body shadow-sm d-flex justify-content-center align-items-center">
                <i class="fal fa-plus-circle fa-5x"></i>
            </div>
        </div>
    </div>
</div>
@endif
@endsection
