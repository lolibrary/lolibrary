FROM amelia/php:7.2 as build

COPY ./api /srv/code
COPY ./models /srv/code/app/Models
COPY ./database /srv/code/database

RUN composer install \
    && php artisan route:cache \
    && rm -rf /var/cache/composer/* \
    && chown -R www-data:www-data storage bootstrap/cache

WORKDIR /srv/code

FROM amelia/php:7.2 as production

COPY --from=build /srv/code /srv/code

RUN apk add --update ca-certificates \
    && update-ca-certificates \
    && composer install --no-dev \
    && chown -R www-data:www-data /srv/code \
    && rm -rf /var/cache/composer/* \
    && rm -rf /var/cache/apk/*
