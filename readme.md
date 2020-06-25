<p align="center"><img height="150" src="/.github/banner.png"></p>
<p align="center">
  <a href="https://patreon.com/lolibrary" title="Support us on Patreon"><img src="/.github/patreon-donate-orange.svg" alt=""></a>
  <a href="https://semaphoreci.com/ameliaikeda/lolibrary" title="Build Status"><img src="https://semaphoreci.com/api/v1/ameliaikeda/lolibrary/branches/master/badge.svg" alt=""></a>
  <a href="https://codeclimate.com/github/lolibrary/lolibrary/maintainability" title="Project Maintainability Score on Code Climate"><img src="https://api.codeclimate.com/v1/badges/4d4b0fa8d8f9d80a00a9/maintainability" alt=""></a>
</p>

## Lolibrary

Lolibrary is a lolita fashion archive website. This repository is a monorepo containing code for what will eventually be all of Lolibrary.


## Getting Started

To get started, you'll need to install Docker and the go compiler.
You'll also need kubectl at a minimum, and if you're a Lolibrary core dev, access to our GKE cluster (ask on Discord, and use your @lolibrary.org email).

As a note, this only really supports macOS right now.

First, run `./bin/install`. This will compile all Go tools and stick them into `$GOPATH/bin`. Add this to your path if needed!

💻 On macOS, you can [install Homebrew](https://brew.sh) to get all of these tools.

## Running internal tools

- To communicate with production (or your local environment), see [Flower CLI](./docs/flower.md)
- To deploy, see [Deploying Lolibrary](./docs/deploying.md)
- To connect to CockroachDB, see [Connecting to CockroachDB](./docs/cockroachdb.md)

### HTTPS certificate

Lolibrary's ingress in dev should be running on HTTPS, as that assumption is made everywhere.

There is a self-signed certificate in the `pki` folder of this repository; you should trust this certificate on your development machine.

Then, provided you follow the guidance below, you'll be able to reach:

- lolibrary.test:443 (main user frontend)
- api.lolibrary.test:443 (all services that start with `service.api.`)
- admin.lolibrary.test:443 (admin user frontend)
- image-proxy.lolibrary.test:443 (a proxy to our configured CDN so that you load images locally)

This is only needed when running full tests across the entire platform. In most cases when you test a service in isolation you don't use the ingress, so you can just find it on localhost.

### Hostnames

Add a dns resolver such as dnsmasq to resolve all `.test` domains to localhost. This will be needed! (`setup.sh`, again, does this automatically on Mac).

On windows, you can get away with just adding `lolibrary.test` and related domains to your hosts file pointing to `127.0.0.1`

## Infrastructure Overview

![Infrastructure Overview](https://github.com/ChunCVL/lolibrary/blob/master/InfrastructureAnalysis.png)

## Security Vulnerabilities

If you discover a security vulnerability within this repo, email [engineering@lolibrary.org](mailto:engineering@lolibrary.org). All security vulnerabilities will be promptly addressed.

## License

Lolibrary's code is licenced under the [BSD Licence](https://opensource.org/licenses/BSD-3-Clause).

You may not use any of Lolibrary Inc's logos or trademarks without prior written permission.

Lolibrary Inc is a 501(c)(3) non-profit incorporated in the USA.
