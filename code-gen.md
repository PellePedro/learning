<details>
   <summary>code-gen scripts</summary>
  [code-gen](https://github.com/kubernetes/code-generator)
  <details>
  <summary>hack/update-codegen.sh</summary>

  ```
  #!/bin/bash
  set -e

  PROJECT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/../" && pwd )"
  PACKAGE_BASE=$(head -n 1 ${PROJECT_ROOT}/go.mod | awk '{ print $2}')
  CODGEN_IMAGE_NAME="antrea/codegen:kubernetes-1.21.0-build.1"
  function docker_run() {
    docker pull ${CODGEN_IMAGE_NAME}
    docker run  \
		  -w /go/src/${PACKAGE_BASE} \
		  -v ${PROJECT_ROOT}:/go/src/${PACKAGE_BASE} \
		  "${CODGEN_IMAGE_NAME}" "$@"
  }
  docker_run /go/src/${PACKAGE_BASE}/hack/update-codegen-dockerized.sh "$@"
  ```
  </details>
  
  <summary>hack/update-codgen-dockerized.sh</summary>
  [code-gen](https://github.com/kubernetes/code-generator)
  <details>
  <summary>hack/update-codgen-dockerized.sh</summary>

  ```
    #!/bin/bash
    set -ex
    PROJECT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/../" && pwd )"
    PACKAGE_BASE=$(head -n 1 ${PROJECT_ROOT}/go.mod | awk '{ print $2}')
    
    PACKAGE_BASE=$1
    CLIENT_PATH=pkg/client
    CLIENT_NAME="versioned"

    client-gen --input-base "${PACKAGE_BASE}/pkg" \
              --input "nodeset/v1alpha1" --clientset-path "${PACKAGE_BASE}/${CLIENT_PATH}" \
              --clientset-name "$CLIENT_NAME" --fake-clientset=true -h ./hack/boilerplate.go.txt
  ```
  </details>
