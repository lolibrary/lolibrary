<?php

namespace Tests\Unit;

use Tests\TestCase;
use App\Rules\Username;

class UsernameRuleTest extends TestCase
{
    /**
     * @param $username
     * @dataProvider getValidUsernames
     */
    public function testValidUsernames($username)
    {
        $this->assertTrue((new Username)->passes('username', $username));
    }

    /**
     * Test valid/invalid usernames.
     *
     * @param mixed $username
     * @dataProvider getInvalidUsernames
     */
    public function testInvalidUsernames($username)
    {
        $this->assertFalse((new Username)->passes('username', $username));
    }

    /**
     * Test that our validation message is fine.
     *
     * @covers \App\Rules\Username::message()
     * @return void
     */
    public function testValidationMessage()
    {
        $this->assertEquals(trans('validation.username'), (new Username)->message());
    }

    /**
     * A dataset of valid usernames.
     *
     * @return array
     */
    public function getValidUsernames()
    {
        return collect([
            123,
            '1amelia',
            'amelia',
            'foo-bar-baz',
            'a42',
            'abc----123',
        ])->map(function ($value) {
            return [$value];
        })->all();
    }

    /**
     * A dataset of invalid usernames.
     *
     * @return array
     */
    public function getInvalidUsernames()
    {
        return collect([
            '-bar-baz',
            'ab',
            null,
            '',
            'a',
            'ab*',
            '😂',
            '😉👏🐝🎉',
        ])->map(function ($value) {
            return [$value];
        })->all();
    }
}
