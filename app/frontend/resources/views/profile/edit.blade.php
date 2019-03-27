@extends('layouts.app')


@section('content')
<div class="container">
<form id="nav-profile" method="POST" action="{{ route('profile.save', $user->username) }}">
    @csrf
    <div class="form-group">
        <label for="profile-name">{{ __('Name') }}</label>
        <input type="text" class="form-control" id="profile-name" placeholder="Enter a name" value="{{ old('name') ?? $user->name }}" name="name">
    </div>
    <div class="form-group">
        <label for="profile-username">{{ __('Username') }}</label>
        <input type="text" readonly class="form-control-plaintext text-muted text-monospace" value="{{ $user->username }}">
        <small class="form-text text-muted">
            To change your username,
            <a class="text-info" href="#" data-toggle="tooltip" title="Changing username is not currently supported, sorry!">click here</a>
        </small>
    </div>
    <div class="form-group">
        <label for="profile-email">{{ __('Email Address') }}</label>
        <input type="email" class="form-control" id="profile-email" aria-describedby="emailHelp" placeholder="Enter an email" value="{{ old('email') ?? $user->email }}" name="email">
        <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
    </div>
    <div class="form-group">
        <label for="level">{{ __('Lolibrary Role') }}</label>
            <select id="level" name="level" class="form-control form-control-chosen">
                @foreach (array_keys($roles) as $role)
                    <option value="{{ $role }}" @if ($user->level() == $role) selected @endif>{{ $roles[$role] }}</option>
                @endforeach
            </select>
        </div>

    <div class="row">
        <div class="col-sm-6 offset-sm-3 col-md-4 offset-md-4">
            <button type="submit" class="btn btn-block btn-outline-primary my-4">{{ __('Save') }}</button>
        </div>
    </div>
</form>

                    <form method="POST" action="{{ route('auth.resend') }}">
                        @csrf

                        <div class="form-group row mb-0">
                            <div class="col-md-6 offset-md-4">
                                <button type="submit" class="btn btn-primary">
                                    {{ __('Send Verification Link') }}
                                </button>
                            </div>
                        </div>
                    </form>
                    <form method="POST" action="{{ route('password.email') }}">
                        @csrf

                        <div class="form-group row mb-0">
                            <div class="col-md-6 offset-md-4">
                                <button type="submit" class="btn btn-primary">
                                    {{ __('Send Password Reset Link') }}
                                </button>
                            </div>
                        </div>
                    </form>
</div>
@endsection
