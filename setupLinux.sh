#Working under Debian 9.6.
#Run sh ./setupLinux.sh
#!/usr/bin/env bash

#Base directory
sudo pwd=$(pwd)

#Rename .env.example to .env
sudo set -e

if [ ! -f .env ]; then
    sudo echo "Copying .env.example to .env"
    sudo cp .env.example .env
fi

# first, trust our tls certificate
sudo echo "Adding a certificate to your trust store (pki/certificate.pem)"
sudo echo "⚠️  You may need to enter your password."
sudo apt-get update
sudo apt-get install ca-certificates -y
sudo cp $pwd/pki/certificate.pem --target-directory="/usr/local/share/ca-certificates/certificate.crt"
sudo mv /usr/local/share/ca-certificates/certificate.pem /usr/local/share/ca-certificates/certificate.crt
sudo update-ca-certificates

#Installing DNS
#sudo echo "edit hostfile by adding 127.0.0.1
#sudo echo nano /etc/hosts
echo "Installing dnsmasq via apt-get"
sudo apt-get install dnsmasq -y
sudo service dnsmasq start

#Install Docker
sudo apt-get update
sudo apt-get install apt-transport-https curl gnupg2 software-properties-common -y
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
sudo apt-get update
sudo apt-get install docker-ce -y
sudo docker run hello-world

#Install Docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

#now, docker-compose up and create the initial volumes/files etc
sudo echo "Starting docker services..."
sudo docker-compose up -d

# run install commands; can be run each time.
sudo echo "Installing next"
sudo docker-compose exec www.lolibrary.test composer install

sudo echo "Checking the database is up"
sudo docker-compose exec www.lolibrary.test php artisan wait:db --timeout=15 --sleep=200

status=$?

if [ $status -eq 1 ]; then
    sudo echo "Database timed out";
    exit 1
fi

sudo docker-compose exec www.lolibrary.test php artisan migrate --seed

sudo echo "All done - it may be a little while until the site comes up, because nodejs is actively building the frontend via Laravel Mix."
