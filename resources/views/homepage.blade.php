@extends('layouts.app')

@section('content')
<div class="container">
    <!-- Search Bar -->
    <div class="row">
        <div class="col-sm-1">
            @include('components.filters')
        </div>
        <div class="col-sm-3">
        </div>
    </div>
</div>
@endsection
