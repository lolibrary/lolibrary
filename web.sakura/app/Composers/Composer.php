<?php

namespace App\Composers;

use Illuminate\View\View;

abstract class Composer
{
    /**
     * How long this key should be cached for.
     *
     * @var string
     */
    protected const DURATION = 1440;

    /**
     * Bind data into the view.
     *
     * @param \Illuminate\View\View $view
     * @return void
     */
    public function compose(View $view)
    {
        $view->with($this->name(), $this->data());
    }

    /**
     * Get a list of models from this composer.
     *
     * @return array
     */
    protected function data()
    {
        $default = function () {
            return $this->load();
        };

        try {
            return cache()->remember($this->key(), static::DURATION, $default);
        } catch (Throwable $e) {
            sentry($e);

            return $default();
        }
    }

    /**
     * Get the cache key for this composer.
     *
     * @return string
     */
    protected function key()
    {
        return 'composer:' . $this->name();
    }

    /**
     * The name of this class.
     *
     * @return string
     */
    protected function name()
    {
        return snake_case(class_basename(static::class));
    }

    /**
     * Get models loaded from the database.
     *
     * @return \Illuminate\Database\Eloquent\Collection|\App\Model[]|\Illuminate\Database\Eloquent\Model[]
     */
    abstract protected function load();
}
