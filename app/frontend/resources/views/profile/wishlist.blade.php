@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-md-4">
            <div class="text-center">
                <img src="{{ default_asset() }}" alt="" style="max-height: 150px; max-width: 150px; margin-top: 10px; margin-bottom: 10px" class="img-thumbnail circle">
            </div>
            
            <div class="list-group">
                    <a href="{{ route('profile') }}" class="list-group-item list-group-item-action">
                        <i class="fal fa-fw fa-user"></i>
                        {{ __('Profile') }}
                    </a>
                    <a href="{{ route('wishlist') }}" class="list-group-item list-group-item-action active">
                        <i class="fal fa-fw fa-star"></i>
                        {{ __('Wishlist') }}
                    </a>
                    <a href="{{ route('closet') }}" class="list-group-item list-group-item-action">
                        <i class="fal fa-fw fa-tags"></i>
                        {{ __('Closet') }}
                    </a>
            </div>
        </div>
        <div class="col-md-8">

        </div>
    </div>
</div>
@endsection
