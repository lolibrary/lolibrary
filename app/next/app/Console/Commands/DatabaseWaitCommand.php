<?php

namespace App\Console\Commands;

use DB;
use Illuminate\Console\Command;

class DatabaseWaitCommand extends Command
{
    /**
     * This command's signature.
     * 
     * @var string
     */
    protected $signature = 'db:wait
                            {--timeout=15 : Total seconds to wait until we exit early.}
                            {--sleep=200 : Milliseconds to usleep between attempts to connect.}
                            {--connection= : The connection to wait on.}';

    /**
     * This command's description.
     * 
     * @var string
     */
    protected $description = 'Wait for the given connection to become available.';

    /**
     * If we should terminate our while loop.
     * 
     * @var bool
     */
    protected $terminate = false;

    /**
     * Wait for the DB to become available before
     * 
     * @return int The status code response of this command.
     */
    public function handle()
    {
        $timeout = (int) $this->option('timeout');
        $sleep = (int) $this->option('sleep');

        $this->signal();

        while (true) {
            if ($this->terminate) {
                return 1;
            }

            if ($this->connect()) {
                return 0;
            }

            usleep($sleep);
        }
    }

    /**
     * Install signal handlers for SIGALRM.
     * 
     * @return void
     */
    protected function signal()
    {
        \pcntl_async_signals(true);
        
        \pcntl_signal(SIGALRM, function () {
            $this->terminate = true;
        });

        \pcntl_alarm($timeout);
    }

    /**
     * Try to connect to the database.
     * 
     * @return bool
     */
    protected function connect()
    {
        try {
            $pdo = DB::connection($this->option('connection', null))->getPdo();

            return $pdo !== null;
        } catch (\Throwable $e) {
            return false;
        }
    }
}
