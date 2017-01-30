#! /bin/bash
./gradlew clean buildLinux
export ver=0.1.0
docker build -t simagix/gin-api:$ver -f src/docker/Dockerfile .
docker push simagix/gin-api:$ver
