# Skaffolding develoment in golang
THis is a sample project in golang that demonstarate building for source debugging and multisage container build.
The sample code implements a http where client can send command to list pods in a namespace. 
The app accedd the k8s API server, and we'll attach a debugger to set breakpoits and debug live code.

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

## General
List available make comands
```
make help

help             Display this help.
clean-docker     Clean dangeling docker images
build-image      Build Container Image
push-image       Push Container Image
run              Run container in docker
exec             exec to container started by Docker
create-cluster   Create Kind Cluster
destroy-cluster  Destroy Kind Cluster
load-image       Load Docker image in Kind
deploy           Deploy pod to kinf
undeploy         Undeploy pod
debug-pod        Debug pod in kind with dlv
send-curl        Send Curl command to pod
```

## Start Kind Cluster
```
make create-cluster
```

## Inspect Dockerfile, Build and Push Docker image
Note the build produces two images. One with code and debug tools and one for release
```
$ make build-image

$ docker images
localhost:5000/sample-k8s-client         0.0.1     10fe72045f59   11 minutes ago      40.9MB
localhost:5000/sample-k8s-client-debug   0.0.1     3c8ed9984aa5   11 minutes ago      1.42GB
```

## Deploy Pod, and ensure it's running 
```
make deploy
kubectl get pod
```

## Attach dlv debugger
The sample service is running as main process in the pod with pid 1.<br/>
The dlv is included in the debug image and we can start dlv and attach to the main program with
```
kubectl exec -it my-pod dlv attach 1
alt.
make debug-pod
```
This attached dlv to the service and anters a debug mode.<br/>

### Set breakpoint in the the function that access to the k8s API
```
# set breakpoint
b client.ListPods

# list code
l client.ListPods

# resume execution
c
```
### Trigger the API with curl to hit breakpoint
The pod has a scrip /debug.sh that sends a curl request to list pods in kube-system namespace
```
kubectl exec -it my-pod /debug.sh

alt.
make send-curl
```

[dlv commands](https://github.com/go-delve/delve/tree/master/Documentation/cli#display)





