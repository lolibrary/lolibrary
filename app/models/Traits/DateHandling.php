<?php

namespace App\Models\Traits;

use App\Models\Pivot;
use InvalidArgumentException;
use Illuminate\Support\Carbon;
use Illuminate\Database\Eloquent\Model as Eloquent;

trait DateHandling
{
    /**
     * Return a timestamp as DateTime object, edited to always be in UTC.
     *
     * @param  mixed  $value
     * @return \Illuminate\Support\Carbon
     * @throws \InvalidArgumentException
     */
    protected function asDateTime($value)
    {
        try {
            $date = parent::asDateTime($value);
        } catch (InvalidArgumentException $e) {
            $date = Carbon::parse($value);
        }

        return $date->setTimezone('UTC');
    }

    /**
     * Create a new pivot model.
     *
     * @param \Illuminate\Database\Eloquent\Model $parent
     * @param array $attributes
     * @param string $table
     * @param bool $exists
     *
     * @param null $using
     * @return \App\Pivot
     */
    public function newPivot(Eloquent $parent, array $attributes, $table, $exists, $using = null)
    {
        return $using
            ? $using::fromRawAttributes($parent, $attributes, $table, $exists)
            : Pivot::fromAttributes($parent, $attributes, $table, $exists);
    }

    /**
     * Get the format for database stored dates.
     *
     * @return string
     */
    public function getDateFormat()
    {
        return Carbon::RFC3339;
    }
}
