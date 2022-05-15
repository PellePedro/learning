#!/bin/bash
set -e

if [[ -z ${CI_PIPELINE_ID+x} ]]; then
  CI_PIPELINE_ID=ci-cluster
fi

kind delete cluster --name $CI_PIPELINE_ID
