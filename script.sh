#!/usr/bin/env bash

# *********************************************************************************************************************
# 服务构建与部署脚本

# 服务名
NAME="member"
VERSION="1.0.0"
IMAGE="registry.cn-hangzhou.aliyuncs.com/mszs/$NAME"
STACK_ENV="stack-umq-dev"

cd /drone/$NAME && pwd && ls -ltrh
docker info
echo  "*** 构建镜像并推送到镜像仓库"
docker build -t $IMAGE:latest ./
rm -rf /drone/$NAME
docker push $IMAGE:latest

echo "*** 构建生产版本标签镜像"
docker tag $IMAGE:latest $IMAGE:$VERSION
docker push $IMAGE:$VERSION

echo "*** pull ms-docker"
STACKGIT="https://github.com/mszsgo/ms-docker.git"
if [ ! -d "/drone/ms-docker/" ];then
  echo "新建 git clone $STACKGIT"
  cd /drone
  git clone $STACKGIT
fi
echo "更新 git pull $STACKGIT"
cd /drone/ms-docker && git pull $STACKGIT

echo "*** docker stack deploy -c /drone/ms-docker/$STACK_ENV/micro-api-stack.yml micro "
docker stack deploy -c /drone/ms-docker/$STACK_ENV/micro-api-stack.yml micro

echo "Successfully...... "

