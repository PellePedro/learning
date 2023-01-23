[Vlads Repo](https://github.com/vladimirvivien/go-grpc)
<br/>
[Go devcontainer](https://github.com/qdm12/godevcontainer)
<details>
  <summary>Go Workspaces</summary>
  
  https://go.dev/doc/tutorial/workspaces
  
  Topology
  ```golang
  /top-directory
    /module1
      go.mod
    /module2
      go.mod
  go.work
  ```
  Commands
  ```
  go work init ./module1
  go work use ./module2
  ```
  
</details>


## Installing
```
GO_VERSION="1.19.5"
GOLANG_TAR=linux-amd64.tar.gz
GOLANG_DOWNLOAD_URL=https://dl.google.com/go/go${GO_VERSION}.${GOLANG_TAR}
curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz
sudo tar -C /usr/local -xzf golang.tar.gz
rm golang.tar.gz
sudo ln -sf /usr/local/go/bin/* /usr/local/bin
rm -rf ${GOLANG_TAR}
```

## Delve

[Delve](https://github.com/go-delve/delve/tree/master/Documentation/cli)
```
dlv debug --headless --listen=:2345 --log --api-version=2 ./cmd/worker
    
```

## VScode
[Vscode](https://github.com/golang/vscode-go/wiki/debugging-legacy)<BR/>
[Remote Debugging](https://github.com/golang/vscode-go/wiki/debugging-legacy#remote-debugging)
```
{
    "name": "Launch remote",
    "type": "go",
    "request": "attach",
    "mode": "remote",
    "substitutePath": [
		{
			"from": "/home/ubuntu/git/letsramp/skyramp",
			"to": "/workspacee",
		},
	],
    "port": 2345,
    "host": "127.0.0.1",
    "debugAdapter": "legacy",
    "cwd": "/home/ubuntu/git/letsramp/skyramp",
}
```




<details>
  <summary>Go Plugin</summary>
  [Doc](https://pkg.go.dev/plugin) <br/>
  [Guide](https://www.sobyte.net/post/2022-03/golang-plugin/)
  package main
  import "fmt"
  var V int
  func F() { fmt.Printf("Hello, number %d\n", V) }
  go build -buildmode=plugin -o plugin_name.so cmd/main 
	
  p, err := plugin.Open("plugin_name.so")
  if err != nil {
	panic(err)
  }
  v, err := p.Lookup("V")
  if err != nil {
	panic(err)
  }
  f, err := p.Lookup("F")
  if err != nil {
	panic(err)
  }
  *v.(*int) = 7
  f.(func())() // prints "Hello, number 7"

  ```  
</details>
<details>
  <summary>GRPC</summary>
  
  <details>
  <summary>Makefile</summary>

  ```
  PROTOC_IMAGE ?= takama/protoc:v0.4.15

  PROTOC_RUN = docker run --rm -v $(shell pwd):/contracts -w /contracts $(PROTOC_IMAGE) -I.

  all: cni

  cni: ## generates stubs for protobuf pkg/api/cni/v1beta1/cni.proto 
	  @echo "+ $@"
	  @$(PROTOC_RUN) --go_out=plugins=grpc,paths=source_relative:. \
	 --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
	 --openapiv2_out=logtostderr=true:. pkg/api/cni/v1beta1/cni.proto

  .PHONY: cni
  ```
  </details>
  
  <details>
  <summary>Directory Structure</summary>

  ```
  ./cni/cni.proto
  syntax = "proto3";
  import "google/protobuf/any.proto";

  package antrea_io.antrea.pkg.apis.cni.v1beta1;

  option go_package = "pkg/apis/cni/v1beta1";

  message CniCmdArgs {
    string container_id = 1;
    string netns = 2;
    string ifname = 3;
    string args = 4;
    string path = 5;
    bytes network_configuration = 6;
  }

  message CniCmdRequest {
    CniCmdArgs cni_args = 1;
  }

  enum ErrorCode {
    UNKNOWN = 0;
    INCOMPATIBLE_CNI_VERSION = 1;
    UNSUPPORTED_FIELD = 2;
    UNKNOWN_CONTAINER = 3;
    INVALID_ENVIRONMENT_VARIABLES = 4;
    IO_FAILURE = 5;
    DECODING_FAILURE = 6;
    INVALID_NETWORK_CONFIG = 7;
    TRY_AGAIN_LATER = 11;
    IPAM_FAILURE = 101;
    CONFIG_INTERFACE_FAILURE = 102;
    CHECK_INTERFACE_FAILURE = 103;
    // these errors are not used by the servers, but we declare them here to
    // make sure they are reserved.
    UNKNOWN_RPC_ERROR = 201;
    INCOMPATIBLE_API_VERSION = 202;
  }

  message Error {
    ErrorCode code = 1;
    string message = 2;
    repeated google.protobuf.Any details = 3;
  }

  message CniCmdResponse {
    bytes cni_result = 1;
    Error error = 2;
  }

  service Cni {
    rpc CmdAdd (CniCmdRequest) returns (CniCmdResponse) {
    }

    rpc CmdCheck (CniCmdRequest) returns (CniCmdResponse) {
    }

    rpc CmdDel (CniCmdRequest) returns (CniCmdResponse) {
    }
  }
  
  ```
  </details>
  
</details>


<details>
  <summary>Git Submodules</summary>

  ```
  git init .
  git submodule add git-repo
  git add .gitmodules
  git add .
  git comit -m "Added submodules"
  ```
  
</details>
