#!/bin/bash
set -euxo pipefail
docker build "$PWD" -t localhost:32000/matchfunction:1
docker push localhost:32000/matchfunction:1
