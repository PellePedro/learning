# Bash Hints
---
## Parameter Expansion

```
[[ -v MYENV ]] && echo "MYENV is defined with value $MYENV"
```

## Pattern Matching

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
