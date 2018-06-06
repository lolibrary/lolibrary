<?php

namespace Tests\Feature;

use App\User;
use Tests\TestCase;
use Illuminate\Foundation\Testing\WithoutEvents;
use Illuminate\Foundation\Testing\DatabaseTransactions;

class UserTest extends TestCase
{
    use DatabaseTransactions, WithoutEvents;

    /**
     * Test a user can be queried by email address.
     *
     * @return void
     */
    public function testUserCanBeQueriedByEmail()
    {
        $expected = factory(User::class)->create();

        $actual = User::email($expected->email)->first();

        $this->assertEquals($expected->id, $actual->id);
    }

    /**
     * Test a user can be queried by email address.
     *
     * @return void
     */
    public function testUserFailsToBeFoundWithInvalidEmail()
    {
        $expected = factory(User::class)->create(['email' => 'valid@example.com']);

        $actual = User::email('invalid@example.com')->first();

        $this->assertNull($actual);
    }

    /**
     * Test a user can be queried by email address in a case insensitive manner.
     *
     * @return void
     */
    public function testUserEmailQueryIsCaseInsensitive()
    {
        $expected = factory(User::class)->create(['email' => 'ValidEmail@example.com']);

        $actual = User::email('validemail@example.com')->first();

        $this->assertEquals($expected->id, $actual->id);

        $actual = User::email('validEmail@example.com')->first();

        $this->assertEquals($expected->id, $actual->id);

        $actual = User::email('VALIDEMAIL@EXAMPLE.COM')->first();

        $this->assertEquals($expected->id, $actual->id);
    }
}
