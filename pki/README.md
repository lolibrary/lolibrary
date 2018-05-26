<p align="center"><img height="150" src="https://lolibrary.org/assets/ban1-01.png"></p>

# Developer PKI

This directory exists just so that you can generate certificates for HTTPS on localhost.

The hostnames that are valid for use in dev are:

- `lolibrary.test`
- `*.lolibrary.test`
- `localhost`
- `127.0.0.1`

There is already a generated file to use in this directory, but you can update it using `generate_cert.go`:

    go run "$(go env GOROOT)/src/crypto/tls/generate_cert.go" -host lolibrary.test,\*.lolibrary.test,localhost,127.0.0.1 -ca

You will need to trust `certificate.pem` locally in order to visit `https://lolibrary.test` depending on if you've changed your docker-compose file.

On macOS: `sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain certificate.pem`

On Windows, use `certmgr.msc` in `Start > Run`.
