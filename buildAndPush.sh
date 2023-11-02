#!/bin/bash
docker build --no-cache -t avaninath/robot-app:manifest-arm64 --platform linux/arm64 .
docker push avaninath/robot-app:manifest-arm64

docker build --no-cache -t avaninath/robot-app:manifest-amd64 --platform linux/amd64 .
docker push avaninath/robot-app:manifest-amd64

#Remove no-cache
docker manifest rm avaninath/robot-app:manifest-multi-arch
docker manifest create avaninath/robot-app:manifest-multi-arch --amend avaninath/robot-app:manifest-amd64 --amend avaninath/robot-app:manifest-arm64
docker manifest push avaninath/robot-app:manifest-multi-arch
