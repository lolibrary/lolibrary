#Working under Debian 9.6 x64 (Tested in VirtualBox v5.2.22r 126460 (Qt5.6.2)
#!/usr/bin/env bash

#---Main Procedure can be found at the end---

#Declare Interfaces.
declare -f startup
declare -f install_sudo
declare -f clone_lolibrary
declare -f configure_everything_for_lolibrary
declare -f start_lolibrary_containers
declare -f ping_lolibrary_website
declare -f menu_lolibrary

#Global variable.
declare -g IS_SUDO_INSTALLED=false
declare -g RUNNING_WITHIN_ROOT_REPOSITORY=false

#Check if sudo is installed.
function startup {
    #Check sudo
    dpkg -s sudo
    
    status_1=$?
    
    # Status 0 OK ->  Installed
    if [ $status_1 -eq 0 ]; then
        IS_SUDO_INSTALLED=true
        return 0
    fi
    
    #Check directory where the script is running from.
    
    #Local Variable
    baseDirectory=$(pwd)
    REPOSITORY_DIRECTORY_NAME="$baseDirectory/lolibrary"
    
    if [ -d "$REPOSITORY_DIRECTORY_NAME" ]; then
    baseDirectory=$(pwd)
    REPOSITORY_DIRECTORY_NAME="$baseDirectory/lolibrary"
    RUNNING_WITHIN_ROOT_REPOSITORY=true
    fi
}

#Install sudo and add user to sudo group.
function install_sudo {
    echo "Su is called to login as root account, please enter your root passsword and use the below commands to configure sudo."
    echo "Install command: apt-get install sudo"
    echo "Add useracount to sudo group: usermod -aG sudo <username>"
    echo "A reboot is required to apply sudo rights the the inputed username. Command: reboot"
    su
}

#Copy last repository of Lolibrary.
#@GIT_LOLIBRARY_URL
function clone_lolibrary {
    #Constants.
    declare -r GIT_LOLIBRARY_URL="https://github.com/lolibrary/lolibrary.git"
    
    #Variables.
    thisScriptName=`basename "$0"`
    baseDirectory=$(pwd)
    REPOSITORY_DIRECTORY_NAME="$baseDirectory/lolibrary"
    
    #Procedure.
    if [ -d "$REPOSITORY_DIRECTORY_NAME" ]; then
        sudo echo "$REPOSITORY_DIRECTORY_NAME already exists"
        read -r -p "Are you sure you want to delete this directory? [y/N] " response
        case "$response" in
        [yY][eE][sS]|[yY])
           sudo rm -rf $REPOSITORY_DIRECTORY_NAME
           sudo echo "Folder deleted"
        ;;
    *)
        return 0
        ;;
    esac
    fi

    sudo apt-get install git -y
    sudo git clone $GIT_LOLIBRARY_URL
    sudo cp ./$thisScriptName $baseDirectory/lolibrary/
    
    return 0
}

#Install the required sofware and configurate software to run the lolibrary website.
#@CERTIFICATE_PATH_REPOSITORY
#@DOCKER_COMPOSER_VERSION
function configure_everything_for_lolibrary {
    #Base directory
    baseDirectory=$(pwd)
    
    #Constants
    declare -r CERTIFICATE_PATH_REPOSITORY="$baseDirectory/pki/certificate.pem"
    declare -r DOCKER_COMPOSER_VERSION="1.23.2"

    #Rename .env.example to .env
    sudo set -e

    if [ ! -f .env ]; then
        sudo echo "Copying .env.example to .env"
        sudo cp .env.example .env
    fi

    # Add TLS certificate.
    sudo echo "Adding a certificate to your trust store (pki/certificate.pem)"
    sudo apt-get update
    sudo apt-get install ca-certificates -y
    sudo cp $CERTIFICATE_PATH_REPOSITORY /usr/local/share/ca-certificates/certificate.crt
    sudo update-ca-certificates

    #Installing DNS.
    sudo echo "Installing dnsmasq via apt-get"
    sudo apt-get install dnsmasq -y
    sudo service dnsmasq start
    sudo cp /etc/host /etc/hosts_backup
    
    #Add hostname to hosts file.
    sed -i "2i127.0.0.1  lolibrary.test lolibrary" /etc/hosts

    #Install Docker.
    sudo apt-get update
    sudo apt-get install apt-transport-https curl gnupg2 software-properties-common -y
    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add
    sudo apt-key fingerprint 0EBFCD88
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
    sudo apt-get update
    sudo apt-get install docker-ce -y
    
    #Test Docker.
    sudo docker run hello-world

    #Install Docker-compose.
    sudo curl -L "https://github.com/docker/compose/releases/download/$DOCKER_COMPOSER_VERSION/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
    sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
    
    #Test Docker Compose.
    sudo docker-compose --version
    
    return 0
}

#Start Composing Lolibrary Containers.
function start_lolibrary_containers {
    #Run correctly please.
    sudo echo "Please run this from lolibrary root folder where docker-compose.yml is located."

    #now, docker-compose up and create the initial volumes/files etc.
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
    
    return 0
}

#Ping the lolibrary website to test if the DNS is correctly working.
#@AMOUNT_PINGS
function ping_lolibrary_website {
    #Constants.
    declare -ir AMOUNT_PINGS=4
    
    ping -c $AMOUNT_PINGS lolibrary.test
    
    status=$?
    
    if [ $status -eq 0 ]; then
        sudo echo "Website Up";
        return 0
    else
        sudo "Website Down"
        sudo "Check DNS program status if hostname is readed, if required please edit /etc/hosts"
        sudo service dnsmasq status
        cat /etc/hosts
        exit 1
    fi
}

#Main menu
#@IS_SUDO_INSTALLED
#@RUNNING_WITHIN_ROOT_REPOSITORY
function menu_lolibrary {
while :
do
    clear
    cat<<EOF
    ==============================
    Menu Lolibrary Installation Linux
    ------------------------------
    Sudo installed: $IS_SUDO_INSTALLED
    Running within root folder of repository: $RUNNING_WITHIN_ROOT_REPOSITORY
    Option 2 and 3 have to be runned within lolibrary root folder 
    where docker-compose.yml is located.
    
    Please enter your choice:
    Install sudo (if not installed)    (1)
    Clone Lolibrary with Git           (2)
    Configure everything for Lolibrary (3)
    Start Lolibrary Containers         (4)
    Test Website Connection Lolibrary  (5)
    Quit                               (Q)
    ------------------------------
EOF
    read -n1 -s
    case "$REPLY" in
    "1")  echo "you chose choice 1"
    install_sudo
    ;;   
    "2")  echo "you chose choice 1"
    clone_lolibrary
    ;;
    "3")  echo "you chose choice 2"
    configure_everything_for_lolibrary
    ;;
    "4")  echo "you chose choice 3"
    start_lolibrary_containers
    ;;
    "5")  echo "you chose choice 4"
    ping_lolibrary_website
    ;;
    "Q")  exit
    ;;
    "q")  exit
    ;; 
     * )  echo "Invalid option."
    ;;
    esac
    sleep 1
done
}

#Main
startup
menu_lolibrary
