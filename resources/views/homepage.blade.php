@extends('layouts.app')

@section('content')
<div class="container">
    <!-- Search Bar -->
    <div class="row">
        <div class="col-sm-5 col-md-4 col-lg-3">
            @include('components.filters')
        </div>
        <div class="col">
            <div class="card">
                <div class="card-body">
                    <input class="form-control form-control-lg form-control-search"
                            id="search"
                            placeholder="{{ __('Start typing to filter') }}"
                            role="search">
                </div>
            </div>

            <div class="row">
                
            </div>
        </div>
    </div>

    <script type="text/template">
        @include('items.item')
    </script>
</div>
@endsection
