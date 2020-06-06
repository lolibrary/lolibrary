<?php

namespace App\Models;

use App\Models\Traits\HasUuid;
use App\Models\Traits\Collection;
use Illuminate\Support\Str;
use App\Models\Traits\DateHandling;
use Illuminate\Database\Eloquent\Model as Eloquent;

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
    use HasUuid, DateHandling;

    /**
     * The namespace UUID used for {@see uuid5()}.
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
     * Remove all guarding from models.
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
     * Create a new Eloquent Collection instance.
     *
     * @param  array  $models
     * @return \App\Models\Collection
     */
    public function newCollection(array $models = [])
    {
        return new Collection($models);
    }
}
