
```
#!/bin/bash

terminal 1.
skyramp up --local-images <target>

terminal 2

export KUBECONFIG=~/.skyramp/workload-config
ns=$(kubectl get ns | grep skyramp-client | awk '{print $1}') 
kubectl get all -n $ns -o yaml >> deployment.log
kubectl describe all -n $ns  >> deployment.log
kubectl get ClusterRole,ClusterRoleBinding  >> deployment.log
kubectl get ClusterRole,ClusterRoleBinding -o yaml  >> deployment.log
