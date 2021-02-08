# Secret Receiver

![Go Test](https://github.com/betorvs/secretreceiver/workflows/Go%20Test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/betorvs/secretreceiver/badge.svg?branch=master)](https://coveralls.io/github/betorvs/secretreceiver?branch=master)

This microservice was created to act as a CRUD for secret in Kubernetes. To be used to automated secret renew in Kubernetes.


# Environment variables

*ENCODING_REQUEST* is used to accepted only encoded requests.

## Others Environment variables

```sh
export PORT=8080

export APP_NAME=secretreceiver

export LOG_LEVEL=INFO
```

# Deploy it in Kubernetes

Create Secret file:
```sh
kubectl create secret generic secretreceiver --from-literal=encodedrequest=LONGHASH --dry-run -o yaml > k8s-secret.yaml
```

Configure Ingress in k8s-deployment.yaml (line 62) change secretreceiver.example.local to use your own domain.

Deploy it:
```sh
kubectl apply -f k8s-secret.yaml
kubectl apply -f k8s-deployment.yaml
```

# How to create secrets in Secret Receiver

Use [secretpublisher][1] command line tool to send encoded requests to Secret Receiver.

# Development

## Test and coverage

Run the tests

```sh 
TESTRUN=true go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```sh
golangci-lint run
```

## Docker Build

```sh
docker build .
```


# References

## Golang Spell
The project was initialized using [Golang Spell](https://github.com/golangspell/golangspell).

## Architectural Model
The Architectural Model adopted to structure the application is based on The Clean Architecture.
Further details can be found here: [The Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html) and in the Clean Architecture Book.


[1]: [https://github.com/betorvs/secretpublisher]