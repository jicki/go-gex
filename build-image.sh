#!/bin/bash

docker build --target accountapi -t="gex/accountapi" .
docker build --target accountrpc -t="gex/accountrpc" .
docker build --target adminapi -t="gex/adminapi" .
docker build --target matchmq -t="gex/matchmq" .
docker build --target matchrpc -t="gex/matchrpc" .
docker build --target orderapi -t="gex/orderapi" .
docker build --target orderrpc -t="gex/orderrpc" .
docker build --target quoteapi -t="gex/quoteapi" .
docker build --target klinerpc -t="gex/klinerpc" .
