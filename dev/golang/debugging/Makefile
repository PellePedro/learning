REPOSITORY=localhost:5000
IMAGE=sample-k8s-client
TAG=0.0.1
DOCKER_IMAGE=sample-goclient:latest
APPNAME=my-service
POD_NAME=my-controller

DOCKER_IMAGE = ${REPOSITORY}/${IMAGE}-debug:$(TAG)
RELEASE_IMAGE = ${REPOSITORY}/${IMAGE}:$(TAG)


# find module name from go.mod to adjust gopath
module_name := $$(head -n 1 go.mod | awk '{print $$2 }')
build-arg += --build-arg MODULE_NAME=$(module_name)


.PHONY: help
help: ## Display this help.
	@echo $(module_name)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


.PHONY: clean-docker
clean-docker: ## Clean dangeling docker images
		docker rm devcontainer --force
	  docker rm $$(docker ps -q -f 'status=exited')
	  docker rmi $$(docker images -q -f "dangling=true")

# ============= Docker Commands for Container ===================
.PHONY: build-image
build-image: ## Build Container Image
	@echo "Building debug image: $(DOCKER_IMAGE)."
	@DOCKER_BUILDKIT=1 docker build --progress plain --force-rm $(build-arg) --target=builder -t $(DOCKER_IMAGE) .
	@echo "Building releast image: $(RELEASE_IMAGE)."
	@DOCKER_BUILDKIT=1 docker build --progress plain --force-rm $(build-arg) --target=target -t $(RELEASE_IMAGE) .

.PHONY: push-image
push-image: ## Push Container Image
	@docker push ${DOCKER_IMAGE}

.PHONY: run-docker
run:  ## Run container in docker
	@echo "Run: $(DOCKER_IMAGE)."
	@docker run -it --rm --name devcontainer --cap-add=SYS_PTRACE --security-opt seccomp=unconfined -p 8090:8090 $(DOCKER_IMAGE)

.PHONY: exec
exec: ## exec to container started by Docker
	@echo "Attachinig to: $(DOCKER_IMAGE)."
	@docker exec -it devcontainer bash

# ============= K8S Commands ===================

.PHONY: create-cluster
create-cluster: ## Create Kind Cluster
	@./scripts/start-cluster.sh

.PHONY: destroy-cluster
destroy-cluster: ## Destroy Kind Cluster
	@./scripts/destroy-cluster.sh

.PHONY: load-image
load-image: ## Load Docker image in Kind
	@kind load docker-image $(DOCKER_IMAGE)

.PHONY: deploy
deploy: ## Deploy pod to kinf
	@kubectl create -f deployments/rbac.yaml
	@kubectl create -f deployments/pod.yaml

.PHONY: undeploy
undeploy: ## Undeploy pod
	@kubectl delete -f deployments

.PHONY: debug-pod
debug-pod: ## Debug pod in kind with dlv
	@kubectl exec -it my-pod dlv attach 1

.PHONY: send-curl
send-curl: ## Send Curl command to pod
	@kubectl exec -it  my-pod /debug.sh

