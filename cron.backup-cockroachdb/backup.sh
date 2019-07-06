#!/bin/bash

set -e

rst="\033[0m"
grn="\033[0;32m"

database=$1
bucket=$2
serviceaccount=$3
pingurl=$4

if [[ "$database" = "" ]]; then
  echo "⚠️  Database not set."
  echo -e " > Usage: backup <${grn}database${rst}> <${grn}gs://bucket${rst}> <${grn}service-account${rst}>"
  exit 1
fi

if [[ "$bucket" = "" ]]; then
  echo "⚠️  Bucket not set."
  echo -e " > Usage: backup <${grn}database${rst}> <${grn}gs://bucket${rst}> <${grn}service-account${rst}>"
  exit 1
fi

if [[ "$serviceaccount" = "" ]]; then
  echo "⚠️  Service account not set."
  echo -e " > Usage: backup <${grn}database${rst}> <${grn}gs://bucket${rst}> <${grn}service-account${rst}>"
  exit 1
fi

if [[ "$pingurl" = "" ]]; then

fi

current_date=$(date +%Y-%m-%dT%T%z)
filename="${database}-${current_date}.sql"

# activate our GCP service account
echo -e "⬆️  Activating service account ${grn}${serviceaccount}${rst}"
gcloud auth activate-service-account $GCP_SERVICE_ACCOUNT_NAME --key-file=/var/run/secrets/google/key.json
echo "✅  Service account activated"

# dump the file
echo -e "💾 Dumping backup to ${grn}${filename}${rst}"
cockroach dump $database > /tmp/$filename
echo "✅  Backup created successfully"

# upload to GCP
echo -e "⌛️  Uploading backup to ${grn}${bucket}/${filename}${rst}"
gsutil cp /tmp/$filename $bucket/$filename
echo "✅  Uploaded successfully"

rm -rf /tmp/$filename

# ping healthcheck to say we uploaded properly.
# curl ...
echo "⌛️  Running cron healthcheck ping"
curl -fsS --retry=3 $pingurl > /dev/null
echo "✅  Pinged successfully."
