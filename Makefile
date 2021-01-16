PROJECT_NAME := k8s-api-go
REGISTRY_NAME := jrmanes

build:
	docker build . -t $(REGISTRY_NAME)/$(PROJECT_NAME):latest

test:

login:
	docker login --username=${DOCKER_USER} --password=${DOCKER_PASS}

push:
	docker push $(REGISTRY_NAME)/$(PROJECT_NAME):latest