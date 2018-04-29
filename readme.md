<p align="center"><img height="150" src="https://lolibrary.org/assets/ban1-01.png"></p>

## Lolibrary

Lolibrary is a lolita fashion archive website. This repository contains code for the API, which is consumed by the chinese mirror and parts of the main site.


## Getting Started

To get started, you'll need to [install Docker](https://www.docker.com/community-edition). This should be your only real requirement to run Lolibrary's code.

To get started, copy `.env.example` to `.env`; this is your entire config and sets environment variables. The default is enough to get started.

### First-time setup

Next, run these commands in order; they'll only need to be re-run if you run `docker-compose down -v` or delete your `vendor` folder.

```sh
docker-compose run app composer install
docker-compose run app php artisan key:generate
docker-compose run app php artisan migrate:fresh --seed
```

After running the last line, you'll see something like this, which you can use to log in as a developer user:

```
Admin email: admin@example.com
Admin password: ZWrDoq8TKf9RyVf9d4SG1qp4hG1CIel8o3HP2ucXioLmsaNwEY6JicNAEmI0KTTW
```

### General running

Now, to run all code/containers, you'll need to run the following. This is your "normal" command to run to start Lolibrary running. This will start everything in the background; to start in the foreground just omit `-d`. To see logs, run `docker-compose logs` in the same directory as `docker-compose.yml`.

```sh
docker-compose up -d
```

This will start the postgres/redis containers, spin up your web container to serve the application, run the queue workers, and build the frontend assets.

After making changes, your queue workers won't automatically restart. To do this, issue the following command:

```sh
docker-compose run app php artisan horizon:terminate
```

docker-compose will automatically restart the container that was running queue workers.

### HTTPS certificate

Lolibrary in dev should be running on HTTPS, as that assumption is made everywhere. There is a self-signed certificate in the `pki` folder of this repository; you should trust this certificate on your development machine.

**On macOS**

- Open `/pki/certificate.pem` in finder and open it.
- Add it to your `login` keychain.
- Open keychain access (if it's not already there) and open the `lolibrary.test` entry.
- Under "Trust", select "When using this certificate: Always", and exit the pane
- Enter your password

**On Windows**

- Open `certmgr.msc` (Certificate Manager)
- Go to Action > All Tasks > Import
- Import `/pki/certificate.pem` from this repo
- Add it to Trusted Root Certification Authorities and mark it as trusted.

### Hostnames

Add the following as aliases for `127.0.0.1` on your machine (`/etc/hosts` on macOS/linux):

- `lolibrary.test`
- `www.lolibrary.test`
- `api.lolibrary.test`

## Security Vulnerabilities

If you discover a security vulnerability within this repo, email [amelia@lolibrary.org](mailto:amelia@lolibrary.org). All security vulnerabilities will be promptly addressed.

## License

Lolibrary's code is licenced under the [BSD Licence](https://opensource.org/licenses/BSD-3-Clause).

You may not use any of Lolibrary Inc's logos or trademarks without prior written permission.

Lolibrary Inc is a 501(c)(3) non-profit incorporated in the USA.
