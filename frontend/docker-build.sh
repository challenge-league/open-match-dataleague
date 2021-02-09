#!/bin/bash
set -euxo pipefail
docker build -t localhost:32000/frontend:latest .
docker push localhost:32000/frontend:latest
