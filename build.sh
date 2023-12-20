#!/bin/bash

SCRIPT=$(readlink -f "$0")
#获取工作目录
WORKDIR=$(dirname $SCRIPT)

cd $WORKDIR

docker build . -t go-cloud-naitve

docker run -p 8000:8000 go-cloud-native