<?php

namespace App\Models\Traits;

use InvalidArgumentException;

trait HasStatus
{
    /**
     * Get a status map array.
     *
     * @var array
     */
    public function getStatusMap()
    {
        if (isset(static::$statuses)) {
            return static::$statuses;
        }

        throw new InvalidArgumentException('static::$statuses not found on ' . static::class);
    }

    /**
     * An attribute getter for a status array.
     *
     * @return array
     */
    public function getStatusesAttribute()
    {
        return collect($this->getStatusMap())->mapWithKeys(function (int $value, string $status) {
            return [$status => $this->hasRawStatus($value)];
        });
    }

    /**
     * Check if this model has a given status.
     *
     * @param string $status
     * @return bool
     */
    public function hasStatus(string $status)
    {
        $code = $this->getRawStatus($status);

        return $this->hasRawStatus($code);
    }

    /**
     * Check model has a specific status, by code.
     *
     * @return void
     */
    public function hasRawStatus(int $status)
    {
        return ($this->getRawStatusCode() & $status) !== 0;
    }

    /**
     * Gets the status code of a model.
     *
     * @return int
     */
    public function getRawStatusCode()
    {
        return (int) $this->status;
    }

    /**
     * Set the raw status code directly.
     *
     * @param int $status
     * @return void
     */
    public function setRawStatusCode(int $status)
    {
        $this->status = $status;
    }

    /**
     * Get a raw (int) status code.
     *
     * @param string $status
     * @return int
     * @throws \InvalidArgumentException
     */
    public function getRawStatus(string $status)
    {
        if (! $this->statusExists($status)) {
            throw new InvalidArgumentException("Tried to get unknown status '$status' on " . static::class);
        }

        return $this->getStatusMap()[$status];
    }

    /**
     * Check if a given status exists on this model.
     *
     * @param string $status
     * @return bool
     */
    public function statusExists(string $status)
    {
        $map = $this->getStatusMap();

        return array_key_exists($status, $map);
    }

    /**
     * Add a specific status to this model.
     *
     * @param string $status
     * @return void
     */
    public function addStatus(string $status, bool $save = true)
    {
        if ($this->hasStatus($status)) {
            return;
        }

        $new = $this->getRawStatusCode() + $this->getRawStatus($status);

        $this->setRawStatusCode($new);

        if ($save) {
            $this->callModelSave();
        }
    }

    public function addStatuses(array $statuses, bool $save = true)
    {
        foreach ($statuses as $status) {
            $this->addStatus($status, false);
        }

        if ($save) {
            $this->callModelSave();
        }
    }

    /**
     * Call a model save if it's available.
     *
     * @return void
     */
    public function callModelSave()
    {
        if (method_exists($this, 'save')) {
            $this->save();
        }
    }
}
