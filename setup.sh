#!/usr/bin/env bash

set -e

if [ ! -f .env ]; then
    echo "Copying .env.example to .env"
    cp .env.example .env
fi

# first, trust our tls certificate
echo "Adding a certificate to your trust store (pki/certificate.pem)"
echo "⚠️  You may need to enter your password."

sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ./pki/certificate.pem

echo "Installing dnsmasq via homebrew"
brew install dnsmasq

sudo brew services start dnsmasq

echo 'address=/.test/127.0.0.1' > $(brew --prefix)/etc/dnsmasq.conf
sudo bash -c "echo 'nameserver 127.0.0.1' > /etc/resolver/test"

# now, docker-compose up and create the initial volumes/files etc
echo "🐳  Starting docker services..."
docker-compose up -d

# run install commands; can be run each time.
echo "💿  Installing next"
docker-compose exec www.lolibrary.test composer install

echo "⏱  Checking the database is up"
docker-compose exec www.lolibrary.test php artisan wait:db --timeout=15 --sleep=200

status=$?

if [ $status -eq 1 ]; then
    echo "Database timed out 😭";
    exit 1
fi

docker-compose exec www.lolibrary.test php artisan migrate --seed

echo "✅  All done - it may be a little while until the site comes up, because nodejs is actively building the frontend via Laravel Mix."
