

```
#!/bin/bash
set -ex

ns=test-rest-demo
services=("cart-service" "checkout-service" "payment-service" "product-catalog-service")

# Iterate over the array
for service in "${services[@]}"
do
    echo "Processing service: $service"
    kubectl get svc "${service}" -n "${ns}" >> "${service}".yaml
    kubectl get deployment "${service}" -n "${ns}" >> "${service}".yaml
done

```
