@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-sm-12 col-md-4 col-lg-3 mb-2 mb-3">
            @include('components.filters')
        </div>

        <div class="col-sm-12 col-md-8 col-lg-9">
            <div class="row mb-3">
                <div class="col px-2">
                    @include('components.search-bar')
                </div>
            </div>

           <div class="row">
                <div class="col">
                    @include('items.listing', ['items' => App\Models\Item::paginate(24)])
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
