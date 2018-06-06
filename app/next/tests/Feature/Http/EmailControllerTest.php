<?php

namespace Tests\Feature\Http;

use App\User;
use Tests\Feature\TestCase;
use App\Notifications\VerifyEmail;
use Illuminate\Support\Facades\Notification;
use Illuminate\Foundation\Testing\WithoutEvents;

class EmailControllerTest extends TestCase
{
    use WithoutEvents;

    public function testVerifyingUserRoute()
    {
        $user = factory(User::class)->create([
            'email_token' => str_random(64),
        ]);

        $url = route('auth.verify', [$user->email, $user->email_token]);

        $this->assertTrue(
            str_contains($url, [urlencode($user->email), $user->email_token])
        );
    }

    public function testVerifyingUser()
    {
        $user = factory(User::class)->create([
            'email_token' => str_random(64),
        ]);

        $url = route('auth.verify', [$user->email, $user->email_token]);

        $this->get($url)->assertStatus(200);

        $user = $user->fresh();

        $this->assertNull($user->email_token);
        $this->assertTrue($user->verified);
    }

    public function testInvalidVerificationLink()
    {
        $user = factory(User::class)->create([
            'email_token' => str_random(64),
        ]);

        $url = route('auth.verify', [$user->email, str_random(64)]);

        $this->get($url)->assertStatus(404);

        $user = $user->fresh();

        $this->assertNotNull($user->email_token);
        $this->assertFalse($user->verified);
    }

    public function testVerificationLinkWhenVerified()
    {
        $user = factory(User::class)->create([
            'email_token' => null,
        ]);

        $url = route('auth.verify', [$user->email, str_random(64)]);

        $this->get($url)->assertStatus(404);

        $user = $user->fresh();

        $this->assertNull($user->email_token);
        $this->assertTrue($user->verified);
    }

    public function testMissingUserForVerification()
    {
        $url = route('auth.verify', ['foo@example.com', str_random(64)]);

        $this->get($url)->assertStatus(404);
    }

    public function testResendingVerification()
    {
        $url = route('auth.resend');
        $user = factory(User::class)->create(['email_token' => str_random(64)]);

        Notification::fake();

        $this->actingAs($user)->post($url)->assertStatus(200);

        Notification::assertSentTo(
            $user,
            VerifyEmail::class,
            function (VerifyEmail $notification, $channels) use ($user) {
                return $notification->user->id === $user->id;
            }
        );
    }

    public function testResendingWhenAlreadyVerified()
    {
        $url = route('auth.resend');
        $user = factory(User::class)->create(['email_token' => null]);

        Notification::fake();

        $this->actingAs($user)->post($url)->assertStatus(302);

        Notification::assertNotSentTo([$user], VerifyEmail::class);
    }
}
