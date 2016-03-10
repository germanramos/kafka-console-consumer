#!/bin/sh

export GO_UID=`id -u`
export GO_GID=`id -g`
export GO_BUILDER=golang:1.6-alpine

BASENAME=$(basename `pwd`)

: ${GO_PIPELINE_NAME:="$BASENAME"}
: ${GO_PIPELINE_COUNTER:=dev}
: ${GO_STAGE_NAME:=build}
: ${VERSION:="0.1.0"}

if [ "$GO_PIPELINE_COUNTER" == "dev" ]; then
    BASE_PATH=""
else
    BASE_PATH=$(docker inspect r-$HOSTNAME |grep UpperDir |cut -d '"' -f 4)
fi

BUILD_COMMAND="docker run  \
       -v $BASE_PATH$PWD:/go \
       -e GO_UID=$GO_UID \
       -e GO_GID=$GO_GID \
       -e GO_PIPELINE_COUNTER=$GO_PIPELINE_COUNTER \
       --rm \
       --name=${GO_PIPELINE_NAME}-${GO_PIPELINE_COUNTER}-${GO_STAGE_NAME} \
       ${GO_BUILDER} \
       ./build.sh"

echo $BUILD_COMMAND
$BUILD_COMMAND
