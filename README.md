# MS API

## Prerequisites

```
# Enable GO111MODULE
$ export GO111MODULE=on

# Initialize app
$ make init

# Install ffmpeg
$ brew install ffmpeg

# Run this command if there is differeces betreen mod.mod and go.sum
$ go mod tidy
```

Set config file at `config/config.ini`.

## How to develop

```
# Update docs
$ swag init

# Run on local
$ make run

# Build on docker
$ make docker-build

# Run on docker
$ make docker-run

# Run external apps
$ make external-run

# End external apps
$ make external-end

# Help
$ make help
```

## API Docs

http://localhost:8080/docs/api/v1/index.html#/

## Deploy

Deploy to minikube

```
# Start minikube
$ minikube start

# Use local image
$ eval $(minikube docker-env)

# Build docker image
$ docker build -t ms-api .

# Deploy
$ kubectl apply -f deploy/configmap.yaml
$ kubectl apply -f deploy/deployment.yaml
$ kubectl apply -f deploy/service.yaml

# Get all status
$ kubectl get all

# Access to deployed app (Click the displayed url)
$ minikube service ms-api --url
```

minikube common commands

```
$ minikube start
$ minikube status
$ minikube dashboard
$ minikube tunnel
$ minikube stop
```
