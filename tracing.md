
# Find k8s api server
```
fgrep http ~/.skyramp/workload-config
```
Example output
```
    server: https://10.0.1.106:6443
```

# 
```
strace -ttt -o status.log skyramp status
1673176929.738588 connect(3, {sa_family=AF_INET, sin_port=htons(6443), sin_addr=inet_addr("10.0.1.106")}, 16) = -1 EINPROGRESS (Operation now in progress)

```
