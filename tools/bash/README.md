# Bash Hints
---
## Parameter Expansion

```
[[ -v MYENV ]] && echo "MYENV is defined with value $MYENV"
```

## Pattern Matching

### Matching ipv6/ipv6 addresses in an array

```
addresses=( 'fd50:f:a:b::10:ffff', '192.168.10.10' )
for address in "${addresses[@]}"; do
   [[ ${address} =~ (\.) ]] && echo "${address} is ipv4"
   [[ ${address} =~ (:) ]]  && echo "${address} is ipv6"   
done


```

### Split

    fqdn=cl1.lab.redhat.com
    
    # Filter out Base Domain from fqdn
    base_domain="${fqdn#*.}"
    # result labs.redhat.com
    
    # Filter out Cluster Name from fqdn
    cluster_name="${fqdn%%.*}"
    # result cl1

    # Read pull secret from registry assets
    pull_secret=${KSDF_PULL_SECRET:-${registry_root}/auth/pull-secret.json}
    pull_secret=$( < "${pull_secret}" )
