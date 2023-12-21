#!/bin/bash

#声明目录空间
SCRIPT=$(readlink -f "$0")

#获取工作目录
WORKDIR=$(dirname $SCRIPT)

# 进入工作目录
cd $WORKDIR

pre_prefix="pre-"
pre_name="go-cloud-native"

function images() {
    local new_tag=$(echo ${pre_prefix}$(date +'%Y%m%d')-$(docker inspect -f '{{.Config.Image}}' ${pre_name} | wc -l | xargs printf '%02d'))

    echo ${new_tag}
    # 编程成镜像，并设置镜像名称
    docker build . -t ${pre_name}:${new_tag}

    # --privileged=true 给容器开特权，这样 root 用户才是真正的 root 用户，否则 root 用户就是【普通用户】
    docker run -d --name ${pre_name} -p 8000:8000 --privileged=true ${pre_name}:${new_tag}
}

images;
