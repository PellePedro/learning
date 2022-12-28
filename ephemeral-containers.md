

https://devopstales.github.io/home/k8s-ephemerald-pod/


Example: general deployment
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: slim
spec:
  selector:
    matchLabels:
      app: slim
  template:
    metadata:
      labels:
        app: slim
    spec:
      containers:
      - name: app
        image: gcr.io/distroless/python3-debian11
        command:
        - python
        - -m
        - http.server
        - '8080'
EOF
```


```
POD_NAME=$(kubectl get pods -l app=slim -o jsonpath='{.items[0].metadata.name}')

# no bash in the container
$ kubectl exec -it -c app ${POD_NAME} -- bash
error: Internal error occurred: error executing command in container: failed to exec in container: failed to start exec "43d1e91f41310fb1ede9fbab741921091edfe116311f18a3881f90f68d06dc13": OCI runtime exec failed: exec failed: unable to start container process: exec: "bash": executable file not found in $PATH: unknown

# ther is sh in the container but limited tools
$ kubectl exec -it -c app ${POD_NAME} -- sh
$# ps
sh: 3: ps: not found
```


```
$ kubectl debug -it --attach=false -c debugger --image=nicolaka/netshoot ${POD_NAME}

$  kubectl get pod ${POD_NAME}   -o jsonpath='{.spec.ephemeralContainers}' | jq

[
  {
    "image": "nicolaka/netshoot",
    "imagePullPolicy": "Always",
    "name": "debugger",
    "resources": {},
    "stdin": true,
    "terminationMessagePath": "/dev/termination-log",
    "terminationMessagePolicy": "File",
    "tty": true
  }
]

$ kubectl get pod ${POD_NAME}   -o jsonpath='{.status.ephemeralContainerStatuses}' | jq
[
  {
    "containerID": "containerd://c3a58d41f5b007aa1d7c2f6758c0d397428bf1d3575380a0661f34efaab4bb34",
    "image": "docker.io/nicolaka/netshoot:latest",
    "imageID": "docker.io/nicolaka/netshoot@sha256:aeafd567d7f7f1edb5127ec311599bb2b8a9c0fb31d7a53e9cff26af6d29fd4e",
    "lastState": {},
    "name": "debugger",
    "ready": false,
    "restartCount": 0,
    "state": {
      "running": {
        "startedAt": "2022-09-07T08:30:44Z"
      }
    }
  }
]
```


```
$ kubectl attach -it -c debugger ${POD_NAME}

$# netstat -tulpn
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 0.0.0.0:8080            0.0.0.0:*               LISTEN      -

$# wget -O - 127.0.0.1:8080
Connecting to localhost:8080 (127.0.0.1:8080)
writing to stdout
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html>
...
</html>
```


```
$ kubectl patch deployment slim --patch '
spec:
  template:
    spec:
      shareProcessNamespace: true'

# or just start the debug container with hare-processes in the first place
$ kubectl debug -it --attach=false -c debugger --image=nicolaka/netshoot --share-processes ${POD_NAME}
```


```
$ kubectl get pods
NAME                    READY   STATUS        RESTARTS   AGE
slim-5f5ffd5958-b9sgt   1/1     Terminating   0          72m
slim-66475779f5-5c27b   1/1     Running       0          20s

$ kubectl get pods
NAME                    READY   STATUS    RESTARTS   AGE
slim-66475779f5-5c27b   1/1     Running   0          48s

$ POD_NAME=$(kubectl get pods -l app=slim -o jsonpath='{.items[0].metadata.name}')

$ kubectl debug -it --attach=false -c debugger --image=nicolaka/netshoot ${POD_NAME}

$ kubectl attach -it -c debugger ${POD_NAME}

$# ps aux
PID   USER     TIME  COMMAND
    1 65535     0:00 /pause
    7 root      0:00 python -m http.server 8080
   14 root      0:01 zsh
   72 root      0:00 ps aux
```
