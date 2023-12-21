#!/bin/bash

pre_prefix="pre-"

function tagging() {
    git push
    git pull --tags
    local new_tag=$(echo ${pre_prefix}$(date +'%Y%m%d')-$(git tag -l "${pre_prefix}$(date +'%Y%m%d')-*" | wc -l | xargs printf '%02d'))
    echo ${new_tag}
    git tag ${new_tag}
    git push origin $new_tag
}

tagging;
