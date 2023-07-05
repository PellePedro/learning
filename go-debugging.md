
```
FROM golang:1.20

RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /app
COPY . .

# Build the Go application
RUN go build -gcflags="all=-N -l" -o app main.go

EXPOSE 2345

CMD ["dlv", "exec", "./app", "--listen=:2345", "--api-version=2", "--headless=true", "--log"]

```

```
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Remote Debug",
      "type": "go",
      "request": "attach",
      "mode": "remote",
      "remotePath": "/app",
      "port": 2345,
      "host": "localhost",
      "showLog": true,
      "substitutePath": [
        {
          "from": "${workspaceFolder}", "
           to": "/app" },
       ]
    }
  ]
}
```
      "substitutePath": [
        {
          "from": "${workspaceFolder}/05-example-web-application/api-golang",
          "to": "/app"
        }
      ]
```
    "substitutePath": [
      { "from": "${workspaceFolder}", "to": "/path/to/remote/workspace" },
      ...
  ]
```

