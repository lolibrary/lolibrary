@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row mb-3">
        <div class="col text-center">
            <h1 class="h3">Moderation Queue</h1>
        </div>
    </div>

    <div class="body">
        <table class="table">
            <thead>
            <tr>
                <th>Name</th>
                <th>Brand</th>
                <th>Submitter</th>
                <th>Links</th>
            </tr>
            </thead>
            <tbody>
            @foreach ($items as $item)
                <tr>
                    <td>
                        <a href="{{ route('items.show', $item) }}">
                            {{ $item->english_name }}
                        </a>
                    </td>
                    <td>{{ $item->brand->name }}</td>
                    <td>
                        @if (!empty($item->submitter()))
                        {{-- TODO: When user editing is implimented, this should link there --}}
                        <a href="#">
                            {{ $item->submitter()->name or 'Unknown'}}
                        </a>
                        @else
                            {{ __('Unknown') }}
                        @endif

                    </td>
                    <td>
                        <a href="{{ route('items.show', $item) }}" class="btn btn-sm btn-outline-primary">
                            View
                        </a>
                        <a href="{{ route('items.edit', $item) }}" class="btn btn-sm btn-outline-success">
                            Edit
                        </a>
                        @senior
                            <a href="#" onclick="event.preventDefault(); $('#delete-item').submit()" class="btn btn-sm btn-outline-danger">Delete</a>

                            {{ Form::open(['route' => ['items.destroy', $item], 'method' => 'delete', 'id' => 'delete-item']) }}
                            {{ Form::close() }}
                        @endsenior
                    </td>
                </tr>
            @endforeach
            </tbody>
        </table>
    </div>

    <div class="text-center" style="font-family: sans-serif">
        {{ $items->links() }}
    </div>
@endsection
