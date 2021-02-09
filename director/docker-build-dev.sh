#!/bin/bash
set -euxo pipefail
docker build "$PWD" -f Dockerfile.dev -t localhost:32000/director-dev:1
docker push localhost:32000/director-dev:1
