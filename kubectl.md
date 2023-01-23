
```
export KUBECONFIG=~/.skyramp/workload-config
kubectl get ns

ns=$(kubectl get ns | grep skyramp-client | awk '{print $1}') 
pods=$(kubectl get pods -n $ns --output=jsonpath={.items[*].metadata.name})
```
