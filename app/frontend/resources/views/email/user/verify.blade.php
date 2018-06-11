@component('mail::message')
# Verify Your Lolibrary Email

Hey {{ $user->username }},

You'll need to verify your email to use your Lolibrary account.
Click the button below to verify it!

@component('mail::button', [
    'url' => route('auth.verify', ['email' => $user->email, 'token' => $user->email_token]),
])
Verify Email
@endcomponent

Thanks,<br>
Lolibrary
@endcomponent
