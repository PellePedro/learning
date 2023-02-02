```
kubectl describe configmaps -n local-path-storage local-path-config 

Name:         local-path-config
Namespace:    local-path-storage
Labels:       <none>
Annotations:  
Data
====
config.json:
----
{
        "nodePathMap":[
        {
                "node":"DEFAULT_PATH_FOR_NON_LISTED_NODES",
                "paths":["/var/local-path-provisioner"]
        }
        ]
}
Events:  <none>
```

This configuration means that on each node in the cluster the provisioner will use the /var/local-path-provisioner directory to provision new persistent volumes!

Let’s check that out.

In this demonstration we will:

- Bring up a multi node kind cluster with /var/local-path-provisioner mounted from the host
- Apply our sample pvc-test.yaml and create a deployment and pvc.
- Show that the persistent volume is in our shared directory
- Backup the persistent volume configuration
- Modify the persistent volume configuration
- Delete and recreate the kind cluster
- Restore the persistent volume configuration
- Redeploy the app and the pvc and show that the data has been restored.




The important bit there is that we needed to modify the old persistent volume manifest to change the retention policy to Retain or when we apply it. It will be immediately deleted.

We also kept the claim and node affinity information in the manifest.

One of the things we have not addressed is making sure the workload detaches from the storage before deleting the cluster! in some cases your data might be corrupted if you didn’t safely shut the app down before deleting the cluster!

Use Case “Volume Mobility”

For this we are going to use a different storage provider! Our intent is to still provide dynamic creation of pvcs but not to configure the pvcs with node affinity.

Fortunately there is an example implementation in the sigs.k8s.io repo! You can check it out here

For us to use this we need to build it and host it somewhere our kind cluster can access it. We also need a manifest that will deploy and configure it.

I’ve already built and pushed the container to mauilion/hostpath-provisioner:dev

The manifest I built for this example is below




```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraMounts:
  - hostPath: /tmp/hostpath
    containerPath: /tmp/hostpath

```

```
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: test
  labels:
    # insert any desired labels to identify your claim
    app: test
spec:
  storageClassName: standard
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      # The amount of the volume's storage to request
      storage: 2Gi
```

```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: task-pv-volume 
  labels:
    type: local
spec:
  storageClassName: manual 
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce 
  persistentVolumeReclaimPolicy: Retain
  hostPath:
    path: "/bitnami" 
```

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test
  template:
    metadata:
      labels:
        app: test
    spec:
      containers:
      - name: name
        image: nginx
        imagePullPolicy: Always
        ports:
        - containerPort: 80
        volumeMounts:
        - name: volume
          mountPath: /pvc/
      volumes: 
      - name: volume
        persistentVolumeClaim:
          claimName: test
```




