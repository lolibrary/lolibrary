<?php

namespace App\Models;

use App\User;

/**
 * User Access Levels
 *
 * @property bool $banned
 * @property int $level
 *
 * {@see \App\User}
 */
trait AccessLevels
{
    /**
     * Return the user's permission level.
     *
     * @return int
     */
    public function level(): int
    {
        if ($this->banned) {
            return User::BANNED;
        }

        return $this->level;
    }

    /**
     * Check if a user is a developer.
     *
     * Used for guarding sensitive functions,
     *   e.g. debug info and feature flags.
     *
     * @return bool
     */
    public function developer(): bool
    {
        return $this->level() >= User::DEVELOPER;
    }

    /**
     * Check if a user is a moderator (above admin).
     *
     * @return bool
     */
    public function admin(): bool
    {
        return $this->level() >= User::ADMIN;
    }

    /**
     * Check if a user is an admin (senior lolibrarian).
     *
     * @return bool
     */
    public function senior(): bool
    {
        return $this->level() >= User::SENIOR_LOLIBRARIAN;
    }

    /**
     * Check if a user is able to process the moderation queue.
     *
     * Lolibrarians can also suggest edits to Items
     *
     * @return bool
     */
    public function lolibrarian(): bool
    {
        return $this->level() >= User::LOLIBRARIAN;
    }

    /**
     * Check if a user is able to perform basic functions
     *
     * @return bool
     */
    public function junior(): bool
    {
        return $this->level() >= User::JUNIOR_LOLIBRARIAN;
    }

    /**
     * Check a user's access role.
     *
     * @return string
     */
    public function getRoleAttribute()
    {
        switch (true) {
            case $this->developer():
                return 'Developer';
            case $this->admin():
                return 'Administrator';
            case $this->senior():
                return 'Senior Lolibrarian';
            case $this->lolibrarian():
                return 'Lolibrarian';
            case $this->junior():
                return 'Junior Lolibrarian';
            case $this->banned:
                return 'Banned User';
            default:
                return 'Regular User';
        }
    }
}
