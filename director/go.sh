#!/bin/bash
set -euxo pipefail
bash docker-build.sh
kubectl delete -f director.yaml
kubectl create -f director.yaml
