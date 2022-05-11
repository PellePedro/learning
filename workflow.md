# Development Workflow in golang
This guide describes setting up a new golang project with k8s ci/cd, command line debuging with [dlv debugger](https://github.com/go-delve/delve) and multistage container build file. The projects are managed by make.

## Prerequisit on local machine
- golang 1.18
- docker
- kind
- kubectl
```
# For MAC install [homebrew](https://brew.sh)
brew install go
brew install kind
brew install kubectl
```

## Action Plan (project from scratch)
- Setting up git for Project
- Setting up CI/CD with kind k8s cluster
- Setting up go module for project
- Setting up Makefile
- Setting up Multi Stage Docker build file (debug/release)
- Implementing a sample application in go accessing the Kubernetes API
- Creating deployment descritors and using kuztomize
- Compiling Application for debugging/release
- Deploying the application to local k8s cluster
- Attaching code level debugger (dlv) and demonstrate debugging  

## Create new Git repository
```
git init .
```
## Setting up CI/CD with kind k8s cluster
```
mkdir -p scripts/start-cluster.sh
```
### Edit file scripts/start-cluster.sh and paste following content
```
#!/bin/sh
set -o errexit

# create registry container unless it already exists
reg_name='kind-registry'
reg_port='5001'
if [ "$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

# create a cluster with the local registry enabled in containerd
cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  podSubnet: 10.10.0.0/16
nodes:  
- role: control-plane
- role: worker
- role: worker
- role: worker
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:5000"]
EOF
kubectl wait --for=condition=ready pods --namespace=kube-system -l k8s-app=kube-dns


# connect the registry to the cluster network if not already connected
if [ "$(docker inspect -f='{{json .NetworkSettings.Networks.kind}}' "${reg_name}")" = 'null' ]; then
  docker network connect "kind" "${reg_name}"
fi

# Document the local registry
# https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${reg_port}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF
```

```
chmod +x scripts/start-cluster.sh
```


# Developing a K8S Controller in GO

## Create new go module for component

```bash
export MODULE=distributed.edge.vmware.com/oga-controller
go mod init ${MODULE}
```

## Add Common Dependencies
```
go get github.com/sirupsen/logrus
```

## Add Kubernetes API Dependencies
There are three go dependencies required to interact with the API Server.

[Client Go](https://pkg.go.dev/k8s.io/client-go)<BR/>
[k8s API](https://pkg.go.dev/k8s.io/api)<BR/>
[k8s API Machinery](https://pkg.go.dev/k8s.io/apimachinery)<BR/>
```
go get k8s.io/api@latest
go get k8s.io/client-go@latest
go get k8s.io/apimachinery@latest

```

# Add liniting rules for golangci
```yaml
cat <<-EOF > .golangci.yaml
run:
  tests: false
  skip-dirs:
    - pkg/mock/testdata

output:
  sort-results: true
linters:
  # because of go1.18
  disable:
    - gosimple
    - staticcheck
    - structcheck
    - unused
  enable:
    - goconst
    - goheader
    - goimports
    - makezero
    - misspell
    - whitespace
EOF
```

# Create a Makefile
```Makefile
DOCKER_IMAGE=dei.mcse.vmware.com/oga-controller:0.0.1

module_name := $$(head -n 1 go.mod | awk '{print $$2 }')
build-arg+=--build-arg MODULE_NAME=$(module_name) 

.PHONY: help
help: ## Display this help.
	@echo $(module_name)
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


# ----------------------------------------------------------------------------------
# General build
.PHONY: fmt
fmt: ## Run go formatting
	go fmt ./...

.PHONY: vet
vet: ## Run go veting
	go vet ./...

.PHONY: lint
lint: ## Run go linting
	go golangci-lint rung ./...

.PHONY: test
test: ## Run go testing
	go test -v ./...

.PHONY: go-build
go-build: main.go ## Compile for Release
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w' -o main ./cmd/main.go

.PHONY: debug-build
debug-build: main.go ## Compile for Debug
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags="all=-N -l" -o main-debug ./cmd/main.go
	
# ============= Docker ===================
.PHONY: build-image
build-image: ## Build Container Image
	@echo "Building image: $(DOCKER_IMAGE)."
	@docker build --progress plain --force-rm $(build-arg) -t $(DOCKER_IMAGE) .
```


## Create multi-stage docker build file (Dockerfile)
```
# syntax=docker/dockerfile:1.3-labs
FROM golang:1.18.1-alpine as base
COPY --from=qmcgaw/binpot:golangci-lint /bin /usr/local/bin/golangci-lint
COPY --from=qmcgaw/binpot:dlv /bin /usr/local/bin/dlv
RUN apk add build-base git bash curl

ARG MODULE_NAME
WORKDIR /go/src/${MODULE_NAME}

ADD .golangci.yaml .

COPY  .golangci.yaml .
COPY  Makefile .
COPY  go.* .
COPY  cmd cmd
COPY  pkg pkg

RUN go mod download
RUN go fmt ./...
#RUN go vet ./...
RUN golangci-lint run ./...
RUN CGO_ENABLED=0 go build -gcflags="all=-N -l" -o /usr/local/bin/app ./cmd

CMD ["/usr/local/bin/app"]
```





### Awesome Multistage Build Files
[Concurrent Mining](https://github.com/ahmetson/concurrent-mining)<BR/>

### Links
[VS Code Go Debug](https://github.com/mipnw/vscode-go-debug/blob/main/Dockerfile)<BR/>
```
# demonstrate "dlv debug", including build tags
FROM golang:1.15.8-alpine AS dev
RUN apk add --no-cache --update git
RUN git clone https://github.com/go-delve/delve && cd delve && go install github.com/go-delve/delve/cmd/dlv
WORKDIR /go/src/app
COPY main.go .
CMD [ "dlv", "--listen=:2345", "--headless", "--api-version=2", "--log", "--build-flags=-tags=mytag", "debug", "main.go", "--", "-multiplier", "2" ]

# demonstrate "dlv exec"
FROM dev as build
RUN CGO_ENABLED=0 go build -gcflags="-N -l" -tags mytag -o /usr/local/bin/experiment2 main.go
CMD [ "dlv", "--listen=:2345", "--headless", "--api-version=2", "--log", "exec", "/usr/local/bin/experiment2", "--", "-multiplier", "2" ]

# demonstrate "dlv exec" from an image without any source code
# in VSCode your launch configuration must use /go/src/app (see WORKDIR in dev stage) as "substitutePath.to" or "remotePath"
FROM alpine AS deploy
COPY --from=build /go/bin/dlv /usr/local/bin/dlv
COPY --from=build /usr/local/bin/experiment2 /usr/local/bin/experiment2
```

```
FROM golang:1.18.1-alpine as buildbase

RUN apk update \
  && apk add git \
  && apk add gcc \
  && apk add libc-dev \
  && go install github.com/go-delve/delve/cmd/dlv@latest

ARG MODULE_NAME

WORKDIR ${MODULE_NAME}
COPY . ${MODULE_NAME}

RUN go mod download
RUN go build -gcflags="all=-N -l" -o /release/server ./cmd

```

```
FROM golang:1.18.1-alpine as buildbase
COPY --from=qmcgaw/binpot:dlv /bin /usr/local/bin/dlv
WORKDIR ${MODULE_NAME}
COPY go.mod ${MODULE_NAME}
COPY go.sum ${MODULE_NAME}
COPY cmd ${MODULE_NAME}/cmd
COPY pkg ${MODULE_NAME}/pkg

RUN go mod download
# Build with debug info
RUN go build -gcflags="all=-N -l" -o /usr/local/bin/server ./cmd

CMD [/usr/local/bin/server]
```

[](https://github.com/banjintaohua/docker/tree/master/Go/build)
```Dockerfile
# syntax=docker/dockerfile:1.4
FROM golang:${GO_VERSION:-1.18.1-alpine} as stage1
COPY --from=qmcgaw/binpot:dlv /bin /usr/local/bin/dlv
COPY --from=golangci/golangci-lint:v1.45.0-alpine /usr/bin/golangci-lint /usr/bin/golangci-lint

WORKDIR /build
ADD go.mod .
ADD go.sum .
ADD cmd ./cmd
ADD pkg ./pkg
ADD .golangci.yml .

RUN go mod download

golangci-lint run --timeout 10m0s ./...


RUN <<EOF
go build -gcflags="all=-N -l" -o /release/server
cp entrypoint.sh /release/entrypoint.sh
chmod +x /release/entrypoint.sh
EOF

# Stage 3: run server with dlv
FROM golang:${GO_VERSION:-1.17.3-alpine} as runner

WORKDIR /
COPY --from=builder /release /release

# Health check by curl
RUN apk update \
  && apk add curl tzdata \
  && rm -rf /tmp/* /var/cache/apk/*

EXPOSE 2345 8080

# CMD ["/release/dlv", "--listen=:2345", "--headless=true", "--check-go-version=false", "--api-version=2", "--accept-multiclient", "exec", "/release/server"]
ENTRYPOINT ["/release/entrypoint.sh"]

```
entrypoint.sh
```bash
#!/usr/bin/env sh

set -e
health_check() {
  if [ "$(curl -f http://localhost:8080)" -ne 0 ]; then
    echo "Application startup failed. Exiting."
    exit 1
  fi
  echo "Application startup successful."
  return 0
}

case "$1" in
"dev")
  dlv --listen=:2345 --headless=true --check-go-version=false --api-version=2 --accept-multiclient exec /release/server || health_check
  ;;
"run")
  /release/server || health_check
  ;;
"health")
  health_check
  ;;
*)
  exec "$@"
  ;;
esac

```






## Sample Http server

Add file cmd/main.go
```
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	readTimeout  = 5
	writeTimeout = 10
	idleTimeout  = 120
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	returnStatus := http.StatusOK
	w.WriteHeader(returnStatus)
	message := fmt.Sprintf("Hello %s!", r.UserAgent())
	w.Write([]byte(message))
}

func main() {
	serverAddress := ":8080"
	l := log.New(os.Stdout, "sample-srv ", log.LstdFlags|log.Lshortfile)
	m := mux.NewRouter()

	m.HandleFunc("/", indexHandler)

	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
		Handler:      m,
	}

	l.Println("server started")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
```
