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

