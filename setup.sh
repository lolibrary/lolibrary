#!/usr/bin/env bash

set -e

if [ ! -f .env ]; then
    echo "Copying .env.example to .env"
    cp .env.example .env
fi

# first, trust our tls certificate
echo "Adding a certificate to your trust store (pki/certificate.pem)"
echo "You may need to enter your password."

sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ./pki/certificate.pem

# now, docker-compose up and create the initial volumes/files etc
echo "Starting services..."
docker-compose up -d

# once we can, composer install and run migrations.
echo "Running commands..."
docker-compose exec web composer install
docker-compose exec web php artisan migrate

if [ -f lolibrary.sql ]; then
   docker-compose exec postgres psql < lolibrary.sql
fi
