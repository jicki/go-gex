FROM golang:1.21-alpine3.19 AS builder

ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct
ENV SRC_PATH ${GOPATH}/gex
WORKDIR ${SRC_PATH}

COPY app common resource ./
RUN apk update && apk upgrade

RUN set -ex \
    && rm -rf go.mod go.sum \
    && apk add git tar \
    && go mod init gex\
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

