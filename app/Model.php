<?php

namespace App;

use App\Models\HasUuid;
use DateTime;
use DateTimeInterface;
use Illuminate\Database\Eloquent\Model as Eloquent;
use Illuminate\Support\Str;

/**
 * A base model for this application.
 *
 * @property string $id The UUID of this model.
 * @property \Carbon\Carbon $created_at
 * @property \Carbon\Carbon $updated_at
 *
 * @property string $url
 *
 * @method static Model find(string $id)
 * @method static Model findOrFail(string $id)
 * @method static Model|\Illuminate\Database\Query\Builder|\Illuminate\Database\Eloquent\Builder where(string|array $column, string $operator = null, mixed $value = null)
 */
abstract class Model extends Eloquent
{
    use HasUuid;

    /**
     * The namespace UUID used for {@see uuid5()}
     *
     * @var string
     */
    public const NAMESPACE_UUID = '56195dda-e864-11e6-98d9-b980ab05cceb';

    /**
     * Remove auto-incrementing ID handling.
     *
     * @var bool
     */
    public $incrementing = false;

    /**
     * Remove all guarding from models
     *
     * @var bool
     */
    protected static $unguarded = false;

    /**
     * Add timezones to date formats.
     *
     * @var string
     */
    protected $dateFormat = 'Y-m-d H:i:sO';

    /**
     * Add the URL for every item to its array form.
     *
     * @var array
     */
    protected $appends = ['url'];

    /**
     * The number of items to show per page.
     *
     * @var int
     */
    protected $perPage = 24;

    /**
     * Get the route key for the model.
     *
     * @return string
     */
    public function getRouteKeyName()
    {
        return 'slug';
    }

    /**
     * Helper attribute for getting the URL to any model.
     *
     * @return string
     */
    public function getUrlAttribute()
    {
        $route = $this->getRouteShowName();

        return route($route, $this);
    }

    /**
     * Default to the model name, lowercase and plural.
     *
     * @return string
     */
    protected function getRouteShowName()
    {
        $class = class_basename($this);

        return Str::plural(Str::lower($class)) . '.show';
    }

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
