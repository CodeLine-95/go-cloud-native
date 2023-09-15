#!/bin/bash
###
 # @Author: CodeLine
 # @Date: 2023-09-15 22:48:13
 # @LastEditors: p-qiaoshuai p-qiaoshuai@xiaomi.com
 # @LastEditTime: 2023-09-15 22:54:47
 # @FilePath: /go-cloud-native/tag.sh
### 

pre_prefix="pre-"

function mi-tag() {
    git push
    git pull --tags
    local new_tag=$(echo ${pre_prefix}$(date +'%Y%m%d')-$(git tag -l "${pre_prefix}$(date +'%Y%m%d')-*" | wc -l | xargs printf '%02d'))
    echo ${new_tag}
    git tag ${new_tag}
    git push origin $new_tag
}

mi-tag;
