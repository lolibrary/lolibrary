docker-compose -f docker-compose.test.yml run test \
    sh -c 'php artisan wait:db \
        && php artisan wait:redis \
        && php artisan migrate:fresh --seed --force --no-interaction \
        && vendor/bin/phpunit --coverage-clover=coverage.xml'
