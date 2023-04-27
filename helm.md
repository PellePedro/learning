

## Deploy skyramp helm project in existing namespace
## Define includes / mustExclude
## Skyramp Clean to delete skyramp managed resources
## Re-deploy skyramp in existing Namespace
## Jobs/Pods annotated as helm hooks are deployed before other resources
## helm hooks are sequentially deployed according to configuration hook-weight
## helm hooks are deleted according to configuration "helm.sh/hook-delete-policy"


```bash
kubectl delete namespace tca
kubectl create namespace tca
kubectl get namespace tcs --show labels
```