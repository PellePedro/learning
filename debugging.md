<details>
  <summary>Debug Container</summary>
  
  ```
  FROM golang:1.19.4 AS build
  WORKDIR /
  COPY . .
  RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
  RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./app

  FROM alpine:3.17
  WORKDIR /
  COPY . .
  COPY --from=build /go/bin/dlv dlv
  COPY --from=build /app app
  ENTRYPOINT [ "/dlv" , "exec", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "/app"]]
  
  
  Command:
    /dbg/go/bin/dlv
    exec
    --headless
    --continue
    --accept-multiclient
    --listen=:56268
    --api-version=2
    /src/server
  ```
  
</details>



## Skyramp
```

b tester_local.go:63
b pkg/tester/tester.go:619
ls pkg/tester/tester.go:619


b pkg/tester/tester.go:742
ls pkg/tester/tester.go:742

b pkg/tester/thrift.go:136
ls pkg/tester/thrift.go:136

```
