# Development from Beginning to Production

## Table of Contents
- Setting up CI/CD with k8s cluster
- Setting up git for Project
- Setting up go module for project
- Setting up Makefile
- Setting up Multi Stage Docker build file
- Implementing a sample application in go accessing the Kubernetes API
- Creating deployment descritors and using kuztomize
- Compiling Application with debugging information
- Deploying the application to local k8s cluster
- Attaching code level debugger (dlv) for debugging  

# Create new Git repository
```
git init .
```

# Create new go module for component
```bash
export MODULE=distributed.edge.vmware.com/oga-controller
go mod init ${MODULE}
```

## Add Common Dependencies
```
go get github.com/sirupsen/logrus

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

# use default linters and some additional linters for now
linters:
  enable:
    - goconst
    - goheader
    - goimports
    - makezero
    - misspell
    - nilerr
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
