# k8s-api-go
Rest API written in Go, ready to deploy in a Kubernetes cluster

## Description

Rest API written in Go, ready to run in a Kubernetes cluster.

## API Spec

- GET /user/{id}
- POST /user
- DELETE /user/{id}

# How it works

Use *Makefile* to build, deploy and test the API

In Makefile:

- make build: Build and push the image to the registry
- make deploy: Deploy the app to the cluster
- make test: Run test to check that the API works


---
Jose Ramón Mañes