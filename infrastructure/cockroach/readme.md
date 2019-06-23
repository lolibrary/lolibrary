# CockroachDB

This is our main database. It boils down to "enterprise high-availability postgres", since you can use any postgres client to connect.

Installed via:

```
helm install --name cockroachdb -f values.yml stable/cockroachdb
```

You then need to approve certificate requests:

Docs: https://www.cockroachlabs.com/docs/v19.1/orchestrate-a-local-cluster-with-kubernetes.html#step-2-start-cockroachdb

`kubectl get csr`, then go through and approve each pod:

`kubectl certificate approve default.node.cockroachdb-cockroachdb-0`
