#!/bin/bash
#
# Bootstrap a Google Cloud Shell with better kubernetes tooling.

set -e

# first, bin directory
mkdir ~/bin
echo "export PATH=\$PATH:\$HOME/bin" >> ~/.bashrc

# now, google service catalog
curl https://github.com/GoogleCloudPlatform/k8s-service-catalog/releases/download/v1.0.0-beta.4/service-catalog-installer-v1.0.0-beta.4-linux.tgz > sc.tar.gz
tar zxvf sc.tar.gz
mv sc cfssl cfssljson ~/bin/

# now, helm.
curl https://storage.googleapis.com/kubernetes-helm/helm-v2.13.1-linux-amd64.tar.gz > helm.tar.gz
tar zxvf helm.tar.gz
mv linux-amd64/helm linux-amd64/tiller ~/bin

# lolibrary-specific stuff
gcloud config set project lolibrary-180111
