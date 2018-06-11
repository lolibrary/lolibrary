@extends('layouts.error')

@section('error')
<h2>{{ __("Sorry, something on our end broke while loading this page.") }}</h2>
<h3>{{ __("We've logged the error and will be looking into it!") }}</h3>
@endsection
