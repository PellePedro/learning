<details>
  <summary>code-gen</summary>
  [code-gen](https://github.com/kubernetes/code-generator)
  <details>
  <summary>client-gen</summary>

  ```
  #!/bin/bash
  PACKAGE_BASE="github.com/kube-node/nodeset"
  CLIENT_PATH=pkg/client
  CLIENT_NAME="versioned"
  client-gen --input-base "${PACKAGE_BASE}/pkg" \
             --input "nodeset/v1alpha1" --clientset-path "${PACKAGE_BASE}/${CLIENT_PATH}" \
             --clientset-name "$CLIENT_NAME" --fake-clientset=true -h ./hack/boilerplate.go.txt
  ```
  </details>
