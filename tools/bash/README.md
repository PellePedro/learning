<<<<<<< Updated upstream
# Bash Hints
---
## Parameter Expansion

```
[[ -v MYENV ]] && echo "MYENV is defined with value $MYENV"
```

## Pattern Matching

### Matching ipv6/ipv6 addresses in an array

```

# Example 1
addresses=( 'fd50:f:a:b::10:ffff' '192.168.10.10' )
local ip6=()
local ip4=()
for address in "${addresses[@]}"; do
   [[ ${address} =~ (\.) ]] && ip4+=("${address}")
   [[ ${address} =~ (:) ]]  && ip6+=("${address}")
done

# Example 2
mapfile -t array <  <( yq -r ".KSDF[].Fabrics[].NTP[]" conf/fabric_definition.yaml)
declare -a ipv4=()
declare -a ipv6=()
for entry in "${array[@]}"; do
    [[ "${entry}" =~ (\.) ]] && ipv4+=("${entry}")
    [[ "${entry}" =~ (:) ]]  && ipv6+=("${entry}")
done¬
¬
IFS=','¬
echo "${ipv4[*]}"¬
echo "${ipv6[*]}"¬
IFS=' '¬

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
=======
# learning
