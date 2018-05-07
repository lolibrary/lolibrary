#!/usr/bin/env bash

set -e

# tear down old
echo "Removing Old Data"
docker-compose -p lolibrary_testing down -v &> /dev/null

# run migrations.
echo "Running Migrations"
docker-compose -p lolibrary_testing run app php artisan migrate --seed --force --step --no-interaction

# run tests
docker-compose -p lolibrary_testing run app php vendor/bin/phpunit "$@"
result=$?

exit $result
