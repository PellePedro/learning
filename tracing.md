
# Find k8s api server
```
fgrep http ~/.skyramp/workload-config
```
Example output
```
    server: https://10.0.1.106:6443
```

# Tracing
```
strace -ttt -o syscalls.log skyramp status

tcpdump -i eth0 -w network.pcap 'tcp port 6443'
``
