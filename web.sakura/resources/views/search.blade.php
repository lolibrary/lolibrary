@extends('layouts.app', ['title' => 'Search'])

@section('content')
<search-page
  :categories="@include('components.categories')"
  :brands="@include('components.brands')"
  :features="@include('components.features')"
  :tags="@include('components.tags')"
  :colors="@include('components.colors')"
  :years=@include('components.years')
  url="{{ route('search') }}"
  endpoint="{{ route('api.search') }}"
></search-page>
@endsection
