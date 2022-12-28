#!/bin/bash

#docker network create --driver=bridge work-net-dev
docker network create \
  --driver=bridge \
  --subnet=172.30.0.0/16 \
  --ip-range=172.30.5.0/24 \
  --gateway=172.30.5.254 \
  work-net-dev
mkdir  /home/taojun/mongors/data1 -p
mkdir  /home/taojun/mongors/data2
mkdir  /home/taojun/mongors/data3

./startdb-dev.sh

sleep 1
docker-compose -f redis-compose-dev.yml up -d
sleep 10
# add mongo user 防止没成功添加
docker exec mongo1 /scripts/dbuser-init.sh
#docker-compose -f meta-svc.yml up -d

echo -e "\033[41;33m mongo redis finish \033[0m"
