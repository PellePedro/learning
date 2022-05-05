# Create new Git repository
```
git init .
```

# Create new go module for component
```
export MODULE=distributed.edge.vmware.com/oga-controller
go mod init ${MODULE}
```

# Create a Makefile
```
DOCKER_IMAGE=dei.mcse.vmware.com/oga-controller:0.0.1
module_name := $$(head -n 1 go.mod | awk '{print $$2 }')
build-arg+=--build-arg MODULE_NAME=$(module_name)

.PHONY help:
help: ## Display this help.
	@echo $(module_name)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


# ============= Go ==============
.PHONY fmt:
fmt: ##
	go fmt

.PHONY vet:
vet: ##
	go vet

.PHONY lint:
lint: ##
	go lint

.PHONY test:
test:
	go test ./...

# ============= Docker ===================
.PHONY: build-image
build-image: ## Build Container Image
	@echo "Building image: $(DOCKER_IMAGE)."
	@docker build --progress plain --force-rm $(build-arg) -t $(DOCKER_IMAGE) .
```


## Create Dockerfile
```
ARG GO_VERSION
FROM golang:1.17.5 as builder

ARG MODULE_NAME

WORKDIR ${MODULE_NAME}
COPY . ${MODULE_NAME}

RUN go mod download

FROM SCRATCH as production




```
