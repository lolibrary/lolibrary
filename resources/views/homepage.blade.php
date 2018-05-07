@extends('layouts.app')

@section('content')
<div class="container">
    <!-- Search Bar -->
    <div class="row">
        <div class="col-sm-4">
            @include('components.filters')
        </div>
        <div class="col-sm">
        </div>
    </div>
</div>
@endsection
