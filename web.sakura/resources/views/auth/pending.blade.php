@extends('layouts.app')

@section('content')
<div class="d-flex justify-content-center align-items-center text-center" style="height: 100%">
    <div>
        <h1>Verification Needed</h1>

        <p class="lead">You'll need to verify your email before you proceed.</p>

        <p>Not got an email yet? Click the button below to resend it.</p>

        <form action="{{ route('auth.resend') }}" method="post">
            @csrf

            <button type="submit" class="btn btn-outline-primary">Resend Email</button>
        </form>
    </div>
</div>
@endsection
