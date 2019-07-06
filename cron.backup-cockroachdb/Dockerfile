FROM cockroachdb/cockroach:v19.1.2 as cockroach
FROM google/cloud-sdk:slim as gcloud
LABEL maintainer=engineering@lolibrary.org

COPY --from=cockroach /cockroach/cockroach /usr/bin/cockroach
COPY backup.sh /usr/bin/backup

RUN chmod +x /usr/bin/backup
