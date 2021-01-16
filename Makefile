PROJECT_NAME := k8s-api-go
REGISTRY_NAME := jrmanes

all: test build deploy

build:
	docker build . -f Dockerfile -t $(REGISTRY_NAME)/$(PROJECT_NAME):latest
	docker push $(REGISTRY_NAME)/$(PROJECT_NAME):latest

test:
	docker build . -f Dockerfile.test --build-arg CACHE_DATE=$(date)

deploy:
	kubectl apply -f ./infrastructure/

clean:
	kubectl delete -f ./infrastructure/

port_forward:
	kubectl port-forward service/$(PROJECT_NAME)-svc 8080:8080 -n $(PROJECT_NAME)