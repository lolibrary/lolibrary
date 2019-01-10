#Working under Debian 9.6 x64 (Tested in VirtualBox v5.2.22r 126460 (Qt5.6.2)
#!/usr/bin/env bash

#---Define functions used for menu at the end--

#Copy last repository of Lolibrary
function clone_lolibrary {
    thisScriptName=`basename "$0"`
    baseDirectory=$(pwd)
    sudo apt-get install git -y
    sudo git clone https://github.com/lolibrary/lolibrary.git
    cp ./$thisScriptName $baseDirectory/lolibrary/
}

function configure_everything_for_lolibrary {
    #Base directory
    baseDirectory=$(pwd)

    #Rename .env.example to .env
    sudo set -e

    if [ ! -f .env ]; then
        sudo echo "Copying .env.example to .env"
        sudo cp .env.example .env
    fi

    # first, trust our tls certificate
    sudo echo "Adding a certificate to your trust store (pki/certificate.pem)"
    sudo apt-get update
    sudo apt-get install ca-certificates -y
    sudo cp $baseDirectory/pki/certificate.pem --target-directory="/usr/local/share/ca-certificates/certificate.crt"
    sudo mv /usr/local/share/ca-certificates/certificate.pem /usr/local/share/ca-certificates/certificate.crt
    sudo update-ca-certificates

    #Installing DNS
    echo "Installing dnsmasq via apt-get"
    sudo apt-get install dnsmasq -y
    sudo service dnsmasq start
    sudo cp /etc/host /etc/hosts_backup
    
    #Add Hostname to hosts file
    sed -i "2i127.0.0.1  lolibrary.test lolibrary" /etc/hosts

    #Install Docker
    sudo apt-get update
    sudo apt-get install apt-transport-https curl gnupg2 software-properties-common -y
    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add
    sudo apt-key fingerprint 0EBFCD88
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
    sudo apt-get update
    sudo apt-get install docker-ce -y
    
    #Test Docker
    sudo docker run hello-world

    #Install Docker-compose
    sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
    sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
    
    #Test Docker Compose
    sudo docker-compose --version
}

function start_lolibrary_containers {
    #Run correctly please
    sudo echo "Please run this from lolibrary root folder where docker-compose.yml is located."

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
}

#Bash Menu
PS3='Please enter your choice (1/2/3/4): '
sudo echo "Option 2 and 3 have to be runned within lolibrary root folder where docker-compose.yml is located."
options=("Clone Lolibrary with Git" "Configure everything for Lolibrary" "Start Lolibrary Containers" "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "Clone Lolibrary with Git")
            echo "you chose choice 1"
            clone_lolibrary
            ;;
        "Configure everything for Lolibrary")
            echo "you chose choice 2"
            configure_everything_for_lolibrary
            ;;
        "Start Lolibrary Containers")
            echo "you chose choice 3"
            start_lolibrary_containers
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done
