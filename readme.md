<p align="center"><img height="150" src="/.github/banner.png"></p>
<p align="center">
  <a href="https://patreon.com/lolibrary" title="Support us on Patreon"><img src="/.github/patreon-donate-orange.svg" alt=""></a>
  <a href="https://semaphoreci.com/ameliaikeda/lolibrary" title="Build Status"><img src="https://semaphoreci.com/api/v1/ameliaikeda/lolibrary/branches/master/badge.svg" alt=""></a>
  <a href="https://codeclimate.com/github/lolibrary/lolibrary/maintainability" title="Project Maintainability Score on Code Climate"><img src="https://api.codeclimate.com/v1/badges/4d4b0fa8d8f9d80a00a9/maintainability" alt=""></a>
  <a href="https://codeclimate.com/github/lolibrary/lolibrary/test_coverage" title="Test Code Coverage on Code Climate"><img src="https://api.codeclimate.com/v1/badges/4d4b0fa8d8f9d80a00a9/test_coverage" alt=""></a>
</p>

## Lolibrary

Lolibrary is a lolita fashion archive website. This repository contains code for the API, which is consumed by the chinese mirror and parts of the main site.


## Getting Started

To get started, you'll need to [install Docker](https://www.docker.com/community-edition). This should be your only real requirement to run Lolibrary's code.

💻 On macOS, [install Homebrew](https://brew.sh) before you run `bash setup.sh`.

To get started, run `bash setup.sh`; you may be prompted for your password.

⚠️ Windows does not currently work with `setup.sh`; you'll need to run commands manually.

Check out `.env`; this is your entire config and sets environment variables. The default is enough to get started.

### General running

To run all code/containers, you'll need to run the following. This is your "normal" command to run to start Lolibrary running. This will start everything in the background; to start in the foreground just omit `-d`. To see logs, run `docker-compose logs` in the same directory as `docker-compose.yml`.

```sh
docker-compose up -d
```

Setup will have already ran this.

This will start the postgres/redis containers, spin up your web container to serve the application, run the queue workers, and build the frontend assets.

After making changes, your queue workers won't automatically restart. To do this, issue the following command:

```sh
docker-compose run app php artisan horizon:terminate
```

docker-compose will automatically restart the container that was running queue workers.

### HTTPS certificate

Lolibrary in dev should be running on HTTPS, as that assumption is made everywhere. There is a self-signed certificate in the `pki` folder of this repository; you should trust this certificate on your development machine.

See the `pki` folder in this directory for more information; `setup.sh` will automatically add this certificate.

### Hostnames

Add a dns resolver such as dnsmasq to resolve all `.test` domains to localhost. This will be needed! (`setup.sh`, again, does this automatically on mac).

## Security Vulnerabilities

If you discover a security vulnerability within this repo, email [amelia@lolibrary.org](mailto:amelia@lolibrary.org). All security vulnerabilities will be promptly addressed.

## License

Lolibrary's code is licenced under the [BSD Licence](https://opensource.org/licenses/BSD-3-Clause).

You may not use any of Lolibrary Inc's logos or trademarks without prior written permission.

Lolibrary Inc is a 501(c)(3) non-profit incorporated in the USA.
