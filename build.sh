#!/usr/bin/env bash

export APPLICATION_NAME=simple
export REPO=chuy08

# Linux commands
export BUILD_TIME=$(date +"%Y-%m-%d_%T")
export BUILD_VERSION=$(cat ./VERSION)

# Golang build flags
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

rm -rf simple

echo "Build Version: $BUILD_VERSION"

go build -ldflags="-X simple/cmd.BuildVersion=$BUILD_VERSION" .

docker build -t $REPO/$APPLICATION_NAME .
docker tag $REPO/$APPLICATION_NAME $REPO/$APPLICATION_NAME:$BUILD_VERSION
#docker image push $REPO/$APPLICATION_NAME:$BUILD_VERSION
