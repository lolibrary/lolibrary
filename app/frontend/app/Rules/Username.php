<?php

namespace App\Rules;

use Illuminate\Contracts\Validation\Rule;

class Username implements Rule
{
    /**
     * The regex used to validate usernames.
     *
     * @var string
     */
    protected const REGEX = '/^[^-_][0-9a-z_-]+$/u';

    /**
     * Determine if the validation rule passes.
     *
     * @param  string  $attribute
     * @param  mixed  $value
     * @return bool
     */
    public function passes($attribute, $value)
    {
        $value = (string) $value;

        return $this->length($value) && preg_match(static::REGEX, $value) > 0;
    }

    /**
     * Get the validation error message.
     *
     * @return string
     */
    public function message()
    {
        return trans('validation.username');
    }

    /**
     * The length.
     *
     * @param $value
     * @return bool
     */
    protected function length($value)
    {
        $length = mb_strlen($value);

        return $length >= 3 && $length <= 40;
    }
}
