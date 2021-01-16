PROJECT_NAME := k8s-api-go
REGISTRY_NAME := jrmanes

build:
	docker build . -f Dockerfile -t $(REGISTRY_NAME)/$(PROJECT_NAME):latest
	docker push $(REGISTRY_NAME)/$(PROJECT_NAME):latest

test:
	docker build . -f Dockerfile.test --build-arg CACHE_DATE=$(date)

deploy:
	kubectl apply -f ./infrastructure/