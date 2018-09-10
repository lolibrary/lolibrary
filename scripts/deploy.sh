#!/bin/bash

set -e

hash=$1

if [ "$hash" = "" ]; then
    echo "Usage: ./deploy.sh FRONTEND_IMAGE_TAG"
    exit 1
fi

echo "Deploying lolibrary/frontend:$hash to production"

kubectl set image deployment/frontend-production frontend=lolibrary/frontend:$hash && \
kubectl set image deployment/horizon horizon=lolibrary/frontend:$hash && \
kubectl rollout status deployment/frontend-production && \
kubectl rollout status deployment/horizon
