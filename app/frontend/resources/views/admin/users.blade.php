@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row mb-3">
        <div class="col text-center">
            <h1 class="h3">Users</h1>
        </div>
    </div>
    <div class="row">
        <div class="col-md-6 col-md-offset-3">
            {{ Form::open(['route' => 'admin.users', 'method' => 'post']) }}
            <div class="form-group">
                <label for="search" class="control-label">Search for a user (name, username, email)</label>
                <div class="input-group">
                    <input type="text" class="form-control" id="search" name="search">
                    <div class="input-group-btn">
                        <button class="btn btn-lg btn-primary">Search</button>
                    </div>
                </div>
            </div>
            {{ Form::close() }}
        </div>
    </div>
    <div class="body">
        <table class="table">
            <thead>
            <tr>
                <th>Username</th>
                <th>Name</th>
                <th>Role</th>
                <th>Links</th>
            </tr>
            </thead>
            <tbody>
            @foreach ($users as $user)
                <tr>
                    <td>
                        <a href="{{ route('profile') }}">
                            {{ $user->username }}
                        </a>
                    </td>
                    <td>{{ $user->name ?: 'N/A' }}</td>
                    <td>
                        {{ $user->getRoleAttribute() }}

                    </td>
                    <td>
                        <a href="{{ route('profile.edit', $user->username) }}" class="btn btn-sm btn-outline-success">
                            Edit
                        </a>
                        @senior
                            <a href="#" onclick="event.preventDefault(); $('#delete-user').submit()" class="btn btn-sm btn-outline-danger">Delete</a>

                            {{ Form::open(['route' => ['profile.destroy', $user], 'method' => 'delete', 'id' => 'delete-user']) }}
                            {{ Form::close() }}
                        @endsenior
                    </td>
                </tr>
            @endforeach
            </tbody>
        </table>
    </div>

    <div class="text-center" style="font-family: sans-serif">
        {{ $users->links() }}
    </div>
@endsection