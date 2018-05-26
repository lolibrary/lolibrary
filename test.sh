#!/usr/bin/env bash

set -e

# run tests

docker-compose \
    -f docker-compose.test.yml \
    -p testing \
        up \
            --build \
            --renew-anon-volumes \
            --exit-code-from test \
            --abort-on-container-exit \
            --remove-orphans

# record the result
result=$?

# bring the env down
docker-compose \
    -p testing -f docker-compose.test.yml \
        down \
            -v \
            --remove-orphans

exit $result
