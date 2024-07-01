#!/bin/bash
if [ -z ${1} ]; then
  echo "Need version param";
  exit 1;
else
  ver=$1;
fi
builddate="$(date '+%Y-%m-%d_%H:%M:%S')"
echo $builddate
docker build -t kstylehub:$ver --build-arg VER=$ver --build-arg BUILDDATE=$builddate .
docker tag kstylehub:$ver krismanpratama/kstylehub:$ver
docker tag kstylehub:$ver krismanpratama/kstylehub:latest
docker push krismanpratama/kstylehub:$ver
docker push krismanpratama/kstylehub:latest
docker image prune --filter label=stage=builder --force
