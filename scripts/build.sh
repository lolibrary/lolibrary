#!/usr/bin/env bash

TAG=lolibrary/frontend:${TAG_NAME:-latest}

cd $(dirname $0)/..

docker build --target=production \
    -f ./app/frontend/Dockerfile \
    -t $TAG \
    ./app

docker push $TAG
