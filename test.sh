#!/usr/bin/env bash

set -e

# remove everything just in case
echo "Setting up Services"
docker-compose -p lolibrary_testing down -v &> /dev/null

# run migrations.
echo "Running Migrations"
docker-compose -p lolibrary_testing run app php artisan migrate --force --step --no-interaction &> /dev/null

# run tests
docker-compose -p lolibrary_testing run app php vendor/bin/phpunit
result=$?

# clean up!
echo "Cleaning up services"
docker-compose -p lolibrary_testing down -v &> /dev/null

exit $result
