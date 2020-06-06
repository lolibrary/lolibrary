FROM abiosoft/caddy:php

LABEL maintainer="Lolibrary Inc <engineering@lolibrary.org>"

WORKDIR /srv

COPY --chown=www-user:www-user ./cache /srv

COPY ./Caddyfile /etc/Caddyfile

RUN apk add --no-cache php7-cli php7-pdo_pgsql curl php7-intl
