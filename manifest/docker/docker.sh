#!/bin/bash

# This shell is executed before docker build.

#! /bin/bash
cd $(dirname $0)/../..
src_dir=$(pwd)

image=$1
echo "build image:"$image

## build go
# docker run --rm -v ${src_dir}:/src  golang:1.20.6 /bin/sh -c "/src/manifest/docker/setup" ||  exit -1

## build image
if [ -d .git ]
then
    COMMIT_SHA=$(git rev-parse HEAD)
    TAG_NAME=$(git tag --points-at HEAD)
    REPO_NAME="Checkout of $(git remote get-url origin) at $(git describe --alway --dirty)"
    BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
fi

# tag
gitVersion=$(git rev-parse HEAD)
dateTime=$(date "+%Y-%m-%d %H:%M:%S")
docker build \
    --build-arg "COMMIT_SHA=$COMMIT_SHA" \
    --build-arg "REPO_NAME=$REPO_NAME" \
    --build-arg "BRANCH_NAME=$BRANCH_NAME" \
    --build-arg GIT_PULL_V="$gitVersion"  \
    --build-arg BUILD_DATE="$dateTime"  \
    -t ${image} \
    -f ${src_dir}/manifest/docker/Dockerfile ${src_dir}


