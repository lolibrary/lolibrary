<?php

namespace App\Console;

use Illuminate\Console\Command;

abstract class WaitCommand extends Command
{
    /**
     * This command's signature.
     *
     * @var string
     */
    protected $signature = 'wait:#TYPE#
                            {--timeout=15 : Total seconds to wait until we exit early.}
                            {--sleep=200 : Milliseconds to usleep between attempts to connect.}
                            {--connection= : The connection to wait on.}';

    /**
     * This command's description.
     *
     * @var string
     */
    protected $description = 'Wait for the given #TYPE# connection to become available.';

    /**
     * If we should terminate our while loop.
     *
     * @var bool
     */
    protected $terminate = false;

    /**
     * The type to append for this command.
     *
     * @var string
     */
    protected const TYPE = 'invalid';

    /**
     * Wait for the given service to become available.
     *
     * @return int The status code response of this command.
     */
    public function handle()
    {
        $sleep = (int) $this->option('sleep');
        $connection = $this->option('connection', null);

        $this->signal();

        while (true) {
            if ($this->terminate) {
                $this->error(
                    $this->replaceType('Timeout reached for #TYPE#, exiting.')
                );

                return 1;
            }

            if ($this->connect($connection)) {
                $this->comment(
                    $this->replaceType('Connected to #TYPE#.')
                );

                return 0;
            }

            $this->comment(
                $this->replaceType('Failed to connect to #TYPE#, trying again...')
            );

            usleep($sleep * 1000);
        }
    }

    /**
     * Install signal handlers for SIGALRM.
     *
     * @return void
     */
    protected function signal()
    {
        $timeout = (int) $this->option('timeout');

        pcntl_async_signals(true);

        pcntl_signal(SIGALRM, function () {
            $this->terminate = true;
        });

        pcntl_alarm($timeout);
    }

    /**
     * Configure the console command using a fluent definition.
     *
     * @return void
     */
    protected function configureUsingFluentDefinition()
    {
        $this->signature = $this->replaceType($this->signature);

        parent::configureUsingFluentDefinition();
    }

    /**
     * Set this command's description.
     *
     * @param string $description
     * @return void
     */
    public function setDescription($description)
    {
        parent::setDescription($this->replaceType($description));
    }

    /**
     * Replace #TYPE# with static::TYPE in a string.
     *
     * @param string $str
     * @return string
     */
    protected function replaceType(string $str)
    {
        return str_replace('#TYPE#', static::TYPE, $str);
    }

    /**
     * Try to connect to the given database.
     *
     * @param string|null $connection
     * @return bool
     */
    abstract protected function connect(?string $connection): bool;
}
