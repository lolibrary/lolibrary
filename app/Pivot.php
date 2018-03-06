<?php

namespace App;

use DateTime;
use DateTimeInterface;
use Illuminate\Database\Eloquent\Relations\Pivot as BasePivot;

/**
 * A custom pivot model for this application.
 *
 * @property \Carbon\Carbon $created_at
 * @property \Carbon\Carbon $updated_at
 */
class Pivot extends BasePivot
{
    /**
     * Add timezones to date formats.
     *
     * @var string
     */
    protected $dateFormat = 'Y-m-d H:i:sO';

    /**
     * Return a timestamp as DateTime object, edited to always be in UTC.
     *
     * @param  mixed  $value
     * @return \Carbon\Carbon
     */
    protected function asDateTime($value)
    {
        $carbon = parent::asDateTime($value);

        return $carbon->setTimezone('UTC');
    }

    /**
     * Prepare a date for array / JSON serialization.
     *
     * @param  \DateTimeInterface  $date
     * @return string
     */
    protected function serializeDate(DateTimeInterface $date)
    {
        return $date->format(DateTime::ISO8601);
    }
}
