ARG DLV_VERSION=v1.20.1                                                                                                         [0/0]
ARG GOPLS_VERSION=v0.11.0

FROM qmcgaw/binpot:dlv-${DLV_VERSION} AS dlv
FROM qmcgaw/binpot:gopls-${GOPLS_VERSION} AS gopls

FROM golang:1.19.2-alpine3.15 AS build
RUN apk add build-base curl git thrift protoc nnn vim

COPY --from=dlv /bin /go/bin/dlv
COPY --from=gopls /bin /go/bin/gopls

WORKDIR /app
ADD . /app
RUN go mod download

RUN mkdir -p /usr/local/lib/skyramp/idl/grpc
RUN mkdir -p /usr/local/lib/skyramp/idl/thrift

RUN go run ./cmd/airgap

RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# ENV GOPROXY=file://$GOPATH/pkg/mod/cache/download

CMD ["sleep" , "360000"]
