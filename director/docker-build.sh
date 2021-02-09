#!/bin/bash
set -euxo pipefail
docker build "$PWD" -t localhost:32000/director:1
docker push localhost:32000/director:1
