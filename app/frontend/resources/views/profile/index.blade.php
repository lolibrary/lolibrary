@extends('profile.layout', ['title' => 'Profile'])

@section('profile')
<form id="nav-profile" method="POST" action="{{ route('profile') }}">
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
        <label for="profile-password">{{ __('Password') }}</label>
        <input type="password" class="form-control" id="profile-password" placeholder="Password" name="password">
    </div>
    <div class="form-group">
        <label for="profile-password-confirm">{{ __('Password Confirmation') }}</label>
        <input type="password" class="form-control" id="profile-password-confirm" placeholder="Password" name="password_confirmation">
        <small class="form-text text-muted">Leave this blank if you don't want to change your password.</small>
    </div>

    <div class="row">
        <div class="col-sm-6 offset-sm-3 col-md-4 offset-md-4">
            <button type="submit" class="btn btn-block btn-outline-primary my-4">{{ __('Save') }}</button>
        </div>
    </div>
</form>
@endsection
