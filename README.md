secretreceiver
=============

This microservice was created to act as a CRUD for secret in Kubernetes. To be used to automated secret renew in Kubernetes.


# Build

```sh
go build
```

# Environment variables

*ENCODING_REQUEST* is used to accepted only encoded requests. To send requests encoded, use [secretpublisher](githut.com/betorvs/secretpublisher) command line.

# Deploy it in Kubernetes

Create Secret file:
```sh
kubectl create secret generic sensubot --from-literal=encodedrequest=LONGHASH --dry-run -o yaml > k8s-secret.yaml
```

Configure Ingress in k8s-deployment.yaml (line 62) change secretreceiver.example.local to use your own domain.

Deploy it:
```sh
kubectl apply -f k8s-secret.yaml
kubectl apply -f k8s-deployment.yaml
```