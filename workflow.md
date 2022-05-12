# Skaffolding develoment in golang
This guide describes setting up a new golang project with k8s ci/cd, command line debuging with [dlv debugger](https://github.com/go-delve/delve) and multistage container build file. The projects are managed by make.

## Prerequisit on local machine
- golang 1.18
- docker
- kind
- kubectl
```
# For MAC, the prereq is installed with kind [homebrew](https://brew.sh)
brew install go
brew install kind
brew install kubectl
```

## Action Plan (project from scratch)
- Setting up git for Project
- Setting up CI/CD with kind k8s cluster and local registry
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
### Create kind startup script
Edit file scripts/start-cluster.sh and paste following
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

### Create cluster destroy script
Edit file scripts/destroy-cluster.sh and paste following
```
#!/bin/sh
set -o errexit

if [[ -z ${CI_PIPELINE_ID+x} ]]; then
  CI_PIPELINE_ID=ci-cluster
fi

kind delete cluster --name $CI_PIPELINE_ID
```
```
chmod +x scripts/destroy-cluster.sh
```

# Developing a sample K8S pod that list pods in a namespace

## Create new go module for component

```bash
export MODULE=distributed.edge.vmware.com/oga-controller
go mod init ${MODULE}
```

## Add Sample Code

create file cmd/main.go and paste the following
```golang
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"oga/controller/pkg/kube/client"
)

const (
	listeningPort = "8090"
)

func listPods(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // Parses the request body
	if err != nil {
		fmt.Println("Failed to parse post")
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Failed to read namespace parameters from request")); err != nil {
			// make linter happy
		}
		return
	}
	namespace := req.Form.Get("ns")
	clientInfo := client.ListPods(namespace)

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("List of Pod Name and IP in namespace %s\n", namespace))
	for k, v := range clientInfo {
		b.WriteString(fmt.Sprintf("%s %s\n", k, v))
	}

	if _, err := w.Write(b.Bytes()); err != nil {
		fmt.Println("Error writing response")
	}
}

func main() {
	http.HandleFunc("/list-pods", listPods)
	fmt.Printf("Server started on port [%s]", listeningPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", listeningPort), nil)
	if err != nil {
		fmt.Printf("Failed to Listen %v\n", err)
	}
}

```
create file pkg/kube/client/client.go and paste the following
```golang
package client

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	clientset *kubernetes.Clientset
)

func init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

// ListPods returns a map of pod Names and podIP for a given namespace
func ListPods(namespace string) map[string]string {
	m := make(map[string]string)
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Failed to list Pods")
		return m
	}
	for _, pod := range pods.Items {
		podIP := pod.Status.PodIP
		podName := pod.ObjectMeta.Name
		fmt.Printf("Found pod with name [%s] host IP[%s]", podName, podIP)
		m[podName] = podIP
	}
	return m
}
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

```


## Create multi-stage Dockerfile
```Dockerfile
# syntax=docker/dockerfile:1.3-labs
FROM golang:1.18.1-alpine as builder
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
RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o /release/app ./cmd

COPY <<-EOF /debug.sh
#!/bin/bash
curl -d \"ns=kube-system\" -X POST http://localhost:8090/list-pods
EOF
RUN chmod +x /debug.sh

CMD ["/usr/local/bin/app"]

FROM scratch as target
COPY --from=builder /release/app /usr/local/bin/app
CMD ["/usr/local/bin/app"]
```

## Create gitlab cicd
Create file .gitlab-ci.yml
```yaml
stages:          # List of stages for jobs, and their order of execution
  - start-cluster
  - build-deploy

deploy cluster:       # This job runs in the build stage, which runs first.
  stage: start-cluster
  script:
    - ./scripts/start-cluster.sh
    - kubectl get nodes --insecure-skip-tls-verify=true
  after_script:
    - ./scripts/destroy-cluster.sh

build-deploy-pod:   # This job runs in the test stage.
  stage: build-deploy
  script:
    - make build-image
    - make push-image
    - echo "Deploying application..."
    - make deploy
    - kubectl get pods
```

## Fetch golang dependencies
```
go mod tidy
```

## Build image
```
make build-image
```
## Start kind cluster
```
make create-cluster
```
## Push Inage to local docker registry
```
make push-image
```

## Deploy Pod to local kind cluster
```
make deploy
```

## Ensure that pod is running
```
kubectl get pods
```

## Start Debugging session
```
make debug-pod
```

## Start Debugging session
```
make send-curl
```


