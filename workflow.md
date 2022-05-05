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


## Create multi-stage docker build file (Dockerfile)
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