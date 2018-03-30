FROM node:9 as build

COPY . .

RUN npm install && npm run production

FROM lolibrary/php:7.2

COPY . /srv/code

COPY --from=build public/js/app.js public/js/app.js
COPY --from=build public/css/app.css public/css/app.css

RUN composer install --no-dev \
    && php artisan route:cache \
    && php artisan view:clear \
    && php artisan storage:link \
    && rm -rf /var/cache/composer/* \
    && chown -R www-data:www-data storage bootstrap/cache
