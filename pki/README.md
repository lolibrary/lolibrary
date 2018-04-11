# Developer PKI

This directory exists just so that you can generate certificates for HTTPS on localhost.

The hostnames that are valid for use in dev are:

- `lolibrary.test`
- `www.lolibrary.test`
- `api.lolibrary.test`
- `localhost`
- `127.0.0.1`

There is already a generated file to use in this directory, but you can update it using [`cfssl`](https://github.com/cloudflare/cfssl).

The required params are in `csr.json`; just copy the output (converting newlines) into `privkey.pem` and `certificate.pem`:

```
cfssl selfsign lolibrary.test csr.json > response.json
```

```json
{
  "cert": "-----BEGIN CERTIFICATE-----\nMIIDKjCCAhKgAwIBAgIILzsw4TEfZugwDQYJKoZIhvcNAQELBQAwJTELMAkGA1UE\nBhMCVVMxFjAUBgNVBAoTDUxvbGlicmFyeSBJbmMwHhcNMTgwNDExMDk1NzQ5WhcN\nMTgwNzExMTYwMjQ5WjAlMQswCQYDVQQGEwJVUzEWMBQGA1UEChMNTG9saWJyYXJ5\nIEluYzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL7kaVCWahYKd8AM\ntwnxm/BgV/OLbiExArg4pY+rSMxtPf5nfd8GlZCEqeN0MIHsei9Il6NJQV9UB6/R\nUPAgWDcUj7Hhr6GoBEhf+lfE/bib6q4IXKn2Dhbh7W+0UaYsgmFPjIGr0pRzr60A\nf5J1X2ABP7qoYz4YNsR9UdNysQJdvzQ1q3mAIMA+g1eio1WqauYdZ13jYR0O+WTX\neesCPmB5Se2FXm/I0372zwuEJ7A0+TqZSs08zTmC/t+HYMxok3QWJjkfPMsb8bW8\nHiyzu20vSkJ/5BUo0dDjqLTI70YbEfM9WoFbCoGJuI3U5kF60OUCzvsjJzIC3Irc\nN2RMensCAwEAAaNeMFwwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUF\nBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBSqS/9hcnrzeuP5\n6NEGIqm/NYEyaDANBgkqhkiG9w0BAQsFAAOCAQEAtjkhN8c2igETnWFFzxz6MzxZ\nJgWfi6xuqxq1nkik/yLP+eJdRbf8+Us7hkgD4enZ6cQJyiX3FqLUk8UAfq320lYO\n9RcKLnOykI63Qn2dYJGPNxXnFGxj0QVapDXystY0HkgFonnuEy/OIXjSDqHt6IaC\nya76eJJOQnMIsO7gcIjMZoSKyVU8FDsuIP02u83WNFoGWLrDV9Sl7+aT49q0ow1U\nIvqai4Q+rEoxeP/Aq6RwcavPuA+4KDR7HAV5uTUMCwBXdRPbwxaTDneoirA+3wxE\ntwt3iDYF117nXEGfFflj7TJ66sL3XREBfr3VfhGLaybotHGSRAIxrG2SPIWr1Q==\n-----END CERTIFICATE-----\n",
  "csr": "-----BEGIN CERTIFICATE REQUEST-----\nMIICzzCCAbcCAQAwJTELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxvbGlicmFyeSBJ\nbmMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC+5GlQlmoWCnfADLcJ\n8ZvwYFfzi24hMQK4OKWPq0jMbT3+Z33fBpWQhKnjdDCB7HovSJejSUFfVAev0VDw\nIFg3FI+x4a+hqARIX/pXxP24m+quCFyp9g4W4e1vtFGmLIJhT4yBq9KUc6+tAH+S\ndV9gAT+6qGM+GDbEfVHTcrECXb80Nat5gCDAPoNXoqNVqmrmHWdd42EdDvlk13nr\nAj5geUnthV5vyNN+9s8LhCewNPk6mUrNPM05gv7fh2DMaJN0FiY5HzzLG/G1vB4s\ns7ttL0pCf+QVKNHQ46i0yO9GGxHzPVqBWwqBibiN1OZBetDlAs77IycyAtyK3Ddk\nTHp7AgMBAAGgZTBjBgkqhkiG9w0BCQ4xVjBUMFIGA1UdEQRLMEmCDmxvbGlicmFy\neS50ZXN0ghJhcGkubG9saWJyYXJ5LnRlc3SCEnd3dy5sb2xpYnJhcnkudGVzdIIJ\nbG9jYWxob3N0hwR/AAABMA0GCSqGSIb3DQEBCwUAA4IBAQBWP34YxzLNzLOvS5V6\nmsYGgMw6PyPTxarezmamVx9/qmv8Jgh/qsBpwP7u33ed8ezNwAzqoYpsEmDBhlKa\nevx19hKtl1cbbu8KHmpsg2UHv6SIsEUq+k3rlitINvBkWhQwnpDJRN9Z18UjoHL5\nvhlt7nZ4jMFlQ2Tr1ftOTdSdpp6aRQLbOra30Pt7H0DA1C2hx2mJlHkgWJT5IK7y\nBrBILs4t5I+iTx53lD4UBVoT1Y6S+KdG48jJ/yDLcOR3nA68aeYX47dIa86Tismv\na1EROKqw5LYm3jTI8Tvmhb/ktz7hUBfT2HLycylXOJTRvElPQrpW33TKXvagwf1d\nMTKN\n-----END CERTIFICATE REQUEST-----\n",
  "key": "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAvuRpUJZqFgp3wAy3CfGb8GBX84tuITECuDilj6tIzG09/md9\n3waVkISp43Qwgex6L0iXo0lBX1QHr9FQ8CBYNxSPseGvoagESF/6V8T9uJvqrghc\nqfYOFuHtb7RRpiyCYU+MgavSlHOvrQB/knVfYAE/uqhjPhg2xH1R03KxAl2/NDWr\neYAgwD6DV6KjVapq5h1nXeNhHQ75ZNd56wI+YHlJ7YVeb8jTfvbPC4QnsDT5OplK\nzTzNOYL+34dgzGiTdBYmOR88yxvxtbweLLO7bS9KQn/kFSjR0OOotMjvRhsR8z1a\ngVsKgYm4jdTmQXrQ5QLO+yMnMgLcitw3ZEx6ewIDAQABAoIBACFqOkUxD0DsdCna\nD/Bdqr5ZHwwyzARjX+Z/g/uyL/wY41E1LINt06CImHSIUjVr/7dnLaQXnqnhO7c7\nHL/r6YU6xCyPJP9XOeixZFEY2pEGOf5c2FuoBq9avQVguyorgGcoVaAdQ69Y0Nen\nzq0Efr1/enhiIX0tH+klRr2Mxg9zq5SIPxV1wdIqd3KGfIaom5TeUC1z+1dNwrp1\nkxj/+Jebvo/JPmqSiQwLbMsQ5b51T3VvPF1sU3J0xWWDpoJUMGulS89fqs8PkOPm\nWlGSR5HDU1r/YeVtluJ8cdic22tFarkAPtDjxsS21m/81imFgGt5CY6ndIKn/OuM\nfNbBBYECgYEAww0Jst/OhEaRNw+urtra6gRHEP7xZweLfGSDCXguKmYnaZzCQ5/V\nUPUDiupBbOOwv10CSl5KqULJr5pPz6wFJ/FiSeMcnVEesN3j/Lb0hgkfTfWJdIWL\nH6d8/gOf61C3tHb+QoLmrnalJz08ctEQXs5KZ5lD2q4lYGLgQabbqOUCgYEA+oqz\nkRvoE40KRZB6lNsZAR307jVvdbnoTsJgRrTQHXLnVvQdHqeVhtrmNyGrU0VOt04W\nbQyQpfRl8X7bLnL70OoeF9ulNRskLA6q03KcA6vxKmWJm162agmMxhJODthy8PrA\nhqbrzFJa221xZV9f3puvuPk8aqNc9XjStniUP98CgYEAg9DR1yLNLDOMe8uLX/vj\nqfcOF/xTJC/DImPC8qlXeavjwsn/tzfTL40FweGiKXaOwiSXIa8rgcvzBUuh2FUG\nfwHwTMLBWceymBC8vNLf2Z1cnJDtPePqK5BDNl84ugyoubsZBdX1E752ylfl/Cox\niACd4/l+E9FQzRzFNVkQZckCgYBzbsJLAjOMBniMEoQX31aG6Dl4IxCGijAZUX8w\nNJCKcO1bZ7+e5xGCf4qALjgUdqNM22KDvb1LaO4rNQphPUL+P0+8KEvWyvmAfwV/\nxJdTLb5AjWW4OwwKkPnWLIrgViOnGbDomTGdAvivjp4nWaj/FHYC4HpQm5Hx20gi\nIC4VuwKBgE0AjfBuBwMs/bnulMJlhvXnKd+/CAILNbKWfCF+PkACt/YvPFcMOUzP\nSCpQ7MAdN1DM2rcHI5MqOozJQRYoxkMJvCsGlBuY/tOVyFrdZNVwJJnXJtubUflr\nsgXDeERIQSZZNUsSd/3KQDtWlbnJ3DxWvOqJTq0ybD19wTSZQTTK\n-----END RSA PRIVATE KEY-----\n"
}
```


Then, just copy out `cert` and `key` to files:

```
echo $(cat response.json | jq .cert) > certificate.pem
echo $(cat response.json | jq .key) > privkey.pem
```

You will need to trust `certificate.pem` locally in order to visit `https://localhost:4443` or `https://lolibrary.test` depending on if you've changed your docker-compose file.

On macOS: `sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain certificate.pem`
