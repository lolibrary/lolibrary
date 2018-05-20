@component('mail::message')
# Verify Your Lolibrary Email

You'll need to verify your email to use your Lolibrary account.
Click the button below to verify it!

@component('mail::button', ['url' => route('auth.verify', $user)])
Verify Email
@endcomponent

Thanks,<br>
Lolibrary
@endcomponent
