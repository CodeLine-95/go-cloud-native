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
  # 获取历史版本的 tag
  old_tag=$(docker images --filter reference="${pre_name}" --format "{{.Tag}}")
  result=$(echo $old_tag | grep "${pre_prefix}")
  new_tag=$(echo ${pre_prefix}$(date +'%Y%m%d')-00)
  if [[ "$result" != "" ]]; then
    if [ -n "$old_tag" ]; then
      new_tag=$(echo ${pre_prefix}$(date +'%Y%m%d')-$(echo ${old_tag} | wc -l | xargs printf '%02d'))
    fi
  fi
  echo ${new_tag}
  # 打包成镜像，并设置镜像名称
  # --no-cache 不使用 docker 缓存，用于解决同一个 Dockerfile 构建不同版本的镜像导致镜像ID一样的问题
  docker build --no-cache . -t ${pre_name}:${new_tag}
  if [[ "$result" != "" ]]; then
    if [ -n "$old_tag" ]; then
      old_container_name=$(docker ps --filter ancestor=${pre_name}:${old_tag} --format "{{.Names}}")
      # 删除上一个运行的容器
      docker stop ${old_container_name}
      docker rm ${old_container_name}
    fi
  fi
  # 启动容器，指定镜像
  # --privileged=true 给容器开特权，这样 root 用户才是真正的 root 用户，否则 root 用户就是【普通用户】
  new_container_name=$(GetRandNum 5)
  docker run -d --name ${pre_name}-${new_container_name} -p 8000:8000 --privileged=true ${pre_name}:${new_tag}
}

# 生成指定长度的随机字符串
function GetRandNum(){
  src_string=$(echo date +%s%N | md5sum |cut -c 1-9)
  length=$1
  echo /dev/urandom | head -n 10 | md5sum | head -c ${length}
}

images;
