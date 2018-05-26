<?php

namespace App\Console\Commands;

use Redis;
use App\Console\WaitCommand;

class RedisWaitCommand extends WaitCommand
{
    /**
     * The name for this command.
     * 
     * @var string
     */
    protected const TYPE = 'redis';

    /**
     * Try to connect to the redis database.
     * 
     * @param string|null $connection
     * @return bool
     */
    protected function connect(?string $connection): bool
    {
        try {
            Redis::connection($connection)->ping();

            return true;
        } catch (\Throwable $e) {
            return false;
        }
    }
}
