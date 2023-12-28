#!/bin/bash

#声明目录空间
SCRIPT=$(readlink -f "$0")

#获取工作目录
WORKDIR=$(dirname $SCRIPT)

# 进入工作目录
cd $WORKDIR

pre_prefix="v-"
pre_name="go-cloud-native"

# 自定义版本号，通过命令行传参
version=$1

function images() {
  new_tag=$(echo ${pre_prefix}${version})
  echo "tag：" + ${new_tag}
  # 获取历史版本的 tag
  old_tag_all=$(docker images --filter reference="${pre_name}" --format "{{.Tag}}")
  # 如果不是新版本镜像，则忽略打包镜像操作
  new_tag_version=$(docker images --filter reference="${pre_name}:${new_tag}" --format "{{.Tag}}")
  if [ ! -n "$new_tag_version" ]; then
      # 打包成镜像，并设置镜像名称
      # --no-cache 不使用 docker 缓存，用于解决同一个 Dockerfile 构建不同版本的镜像导致镜像ID一样的问题
      docker build --no-cache . -t ${pre_name}:${new_tag}
  fi
  # shellcheck disable=SC2199
  # shellcheck disable=SC2076
  if [[ " ${old_tag_all[@]} " =~ "${pre_prefix}" ]]; then
    if [ -n "$old_tag_all" ]; then
      # shellcheck disable=SC2068
      for old_tag in ${old_tag_all[@]}
      do
        old_container_name=$(docker ps --filter ancestor=${pre_name}:${old_tag} --format "{{.Names}}")
        result=$(echo $old_container_name | grep "${pre_name}")
        if [[ "$result" != "" ]]; then
          # 删除上一个运行的容器
          echo "stop：" + $(docker stop ${old_container_name})
          echo "rm：" + $(docker rm ${old_container_name})
        fi
      done
    fi
  fi
  # 启动容器，指定镜像
  # --privileged=true 给容器开特权，这样 root 用户才是真正的 root 用户，否则 root 用户就是【普通用户】
  new_container_name=$(GetRandNum 5)
  echo "new run name: " + ${pre_name}-${new_container_name}
  echo "run：" + $(docker run -d --name ${pre_name}-${new_container_name} -p 8000:8000 --privileged=true ${pre_name}:${new_tag})
}

# 生成指定长度的随机字符串
function GetRandNum(){
  length=$1
  echo $(seq 0 9 |sort -R |xargs |tr -d ' ' |md5sum |cut -c -${length})
}

images;