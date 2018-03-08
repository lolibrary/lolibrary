#!/usr/bin/env bash

echo "Setting up postgres..."
docker rm --force lolibrary-postgres &> /dev/null
docker run -d -e POSTGRES_USER=lolibrary --name lolibrary-postgres -p 55432:5432 postgres:10 &> /dev/null

echo "Setting up redis..."
docker rm --force lolibrary-redis &> /dev/null
docker run -d --name lolibrary-redis -p 16379:6379 redis &> /dev/null

echo "Waiting for postgres to be ready..."
until docker run --rm --link lolibrary-postgres:pg postgres:10 pg_isready -U lolibrary -h pg &> /dev/null; do sleep 0.5; done

echo "Waiting for redis to be ready..."
until docker run --rm --link lolibrary-redis:ll redis redis-cli -h ll ping &> /dev/null; do sleep 0.5; done

DB_CONNECTION=testing php artisan migrate --force --step

vendor/bin/phpunit

docker rm --force lolibrary-postgres &> /dev/null
docker rm --force lolibrary-redis &> /dev/null
