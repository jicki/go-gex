#!/bin/bash


network_exists=$(docker network ls --format "{{.Name}}" --filter "name=gex")
if [ -z "$network_exists" ]; then
    docker network create gex
    echo "网络 gex 创建成功！"
fi

lang=$(cat ../../resource/language/zh-CN.yaml)

match=$(cat ../../app/match/rpc/etc/match-deploy.yaml)

docker exec -it etcd /usr/local/bin/etcdctl put language/zh-CN -- "$lang"

docker exec -it etcd /usr/local/bin/etcdctl put config/match/BTC_USDT -- "$match"

docker exec -it pulsar /pulsar/bin/pulsar-admin namespaces create public/trade


