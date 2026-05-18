# dummy-deployment

A minimal Go web app that displays the value of the `DISPLAY_TEXT` environment variable in the browser. Built for practicing [Kargo](https://kargo.io/) promotions across stages.

## Run locally

```bash
go run main.go
# or with custom text
DISPLAY_TEXT="staging environment" go run main.go
```

Visit [http://localhost:8080](http://localhost:8080).

## Docker

```bash
docker build -t dummy-deployment .
docker run -p 8080:8080 -e DISPLAY_TEXT="running in docker" dummy-deployment
```

## Deploy to Kubernetes

```bash
kubectl apply -k k8s/
```

This creates:

| Resource | Name | Purpose |
|----------|------|---------|
| Namespace | `dummy-deployment` | Isolates the workload. |
| Deployment | `dummy-deployment` | Runs the app (1 replica) |
| Service | `dummy-deployment` | Exposes port 80 → 8080 |

The `ConfigMap` (`dummy-deployment-config`) that supplies `DISPLAY_TEXT` is managed in the ArgoCD repo, not here.

## Kargo

The Kubernetes manifests use Kustomize and are structured so Kargo can promote config and image changes across stages. Kargo resources are not included — add them when ready.
