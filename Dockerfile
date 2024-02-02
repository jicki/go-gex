FROM golang:1.21-alpine3.19 AS builder

ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
ENV SRC_PATH ${GOPATH}/gex
WORKDIR ${SRC_PATH}

RUN mkdir -p ${SRC_PATH}/{app,common,resource}
COPY app ./app
COPY common ./common
COPY resource ./resource 

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk update && apk upgrade

RUN set -ex \
    && rm -rf go.mod go.sum \
    && apk add git tar \
    && go mod init gex\
    && go mod tidy \
    && go mod download \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/accountapi ./app/account/api/account.go \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/accountrpc ./app/account/rpc/account.go \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/adminapi ./app/admin/api/admin.go \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/matchmq ./app/match/mq/match.go \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/matchrpc ./app/match/rpc/match.go \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/orderapi ./app/order/api/order.go \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/orderrpc ./app/order/rpc/order.go \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/quoteapi ./app/quotes/api/quote.go \
    && CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/klinerpc ./app/quotes/kline/rpc/kline.go


## name is accountapi
FROM debian:stretch-slim AS accountapi

WORKDIR /app

COPY --from=builder /go/gex/bin/accountapi /app/accountapi
COPY --from=builder /go/gex/app/account/api/etc/account-deploy.yaml /app/account.yaml


## name is accountrpc 
FROM debian:stretch-slim AS accountrpc

WORKDIR /app

COPY --from=builder /go/gex/bin/accountrpc /app/accountrpc
COPY --from=builder /go/gex/app/account/rpc/etc/account-deploy.yaml /app/account.yaml

## name is adminapi
FROM debian:stretch-slim AS adminapi 

WORKDIR /app

COPY --from=builder /go/gex/bin/adminapi /app/adminapi
COPY --from=builder /go/gex/app/admin/api/etc/admin-deploy.yaml /app/admin.yaml

## name is klinerpc
FROM debian:stretch-slim AS klinerpc

WORKDIR /app

COPY --from=builder /go/gex/bin/klinerpc /app/klinerpc
COPY --from=builder /go/gex/app/quotes/kline/rpc/etc/kline-deploy.yaml /app/kline.yaml

## name is matchmq
FROM debian:stretch-slim AS matchmq

WORKDIR /app

COPY --from=builder /go/gex/bin/matchmq /app/matchmq
COPY --from=builder /go/gex/app/match/mq/etc/match-deploy.yaml /app/match.yaml

## name is matchrpc
FROM debian:stretch-slim AS matchrpc

WORKDIR /app

COPY --from=builder /go/gex/bin/matchrpc /app/matchrpc
COPY --from=builder /go/gex/app/match/rpc/etc/match-deploy.yaml /app/match.yaml

## name is orderapi
FROM debian:stretch-slim AS orderapi

WORKDIR /app

COPY --from=builder /go/gex/bin/orderapi /app/orderapi
COPY --from=builder /go/gex/app/order/api/etc/order-deploy.yaml /app/order.yaml

## name is orderrpc
FROM debian:stretch-slim AS orderrpc

WORKDIR /app

COPY --from=builder /go/gex/bin/orderrpc /app/orderrpc
COPY --from=builder /go/gex/app/order/rpc/etc/order-deploy.yaml /app/order.yaml

## name is quoteapi
FROM debian:stretch-slim AS quoteapi

WORKDIR /app

COPY --from=builder /go/gex/bin/quoteapi /app/quoteapi
COPY --from=builder /go/gex/app/quotes/api/etc/quote-deploy.yaml /app/quote.yaml

