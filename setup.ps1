#Working under Windows 10.
#Run PowerShell as follow to bypass restriction of downloaded files or limited execution environment.
#PowerShell.exe -ExecutionPolicy Bypass -File ./install.ps1

#Get root directory.
Write-Host "Retrieving your current path."
$baseDirectory = (Get-Location).path

#Check enivroment configuration filename if false then  copy and rename file to '.env'.
if (-Not (Get-Childitem -Path *.env)) { Write-Host "Copying .env.example to .env"; Copy-Item $baseDirectory\.env.example -destination $basedirectory\.env }

#Install certificate to trust the HTTPS server (requires administrator rights).
#Write-Host "Adding certificate of Lolibrary to your keychain."
certutil -addstore -enterprise -f "Root" $baseDirectory/pki/certificate.pem

#Edit your host file and do not forget removing the # and overwrite it with the current host (without extensions).
# 	127.0.0.1       lolibrary.test 
Write-Host "Opening Notepad, to resolve the hostname, please add the following line to the file  	   127.0.0.1       lolibrary.test"
Start-Process -FilePath notepad.exe -Verb runas -ArgumentList "$env:SystemRoot\system32\drivers\etc\hosts"

# Pauze after your are finished.
Read-Host -Prompt "Press Enter to continue"

Write-Host "Starting docker services..."
docker-compose up -d

Write-Host "Installing next."
docker-compose exec www.lolibrary.test composer install

Write-Host "Checking the database is up."
docker-compose exec www.lolibrary.test php artisan wait:db --timeout=15 --sleep=200

# Retrieve status of last command.
$status=$?

If ($status -eq 1) { Write-Host "Database timed out"; exit 1 }

Write-Host "All done - it may be a little while until the site comes up, because nodejs is actively building the frontend via Laravel Mix."

