<?php

namespace App\Console\Commands;

use DB;
use App\Console\WaitCommand;

class DatabaseWaitCommand extends WaitCommand
{
    /**
     * The name for this command.
     *
     * @var string
     */
    protected const TYPE = 'db';

    /**
     * Try to connect to the database.
     *
     * @param string|null $connection
     * @return bool
     */
    protected function connect(?string $connection): bool
    {
        try {
            $pdo = DB::connection($connection)->getPdo();

            return $pdo !== null;
        } catch (\Throwable $e) {
            return false;
        }
    }
}
