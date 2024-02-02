package main

import (
	"flag"
	"gex/app/match/mq/internal/config"
	"gex/app/match/mq/internal/consumer"
	"gex/app/match/mq/internal/logic"
	"gex/app/match/mq/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "app/match/mq/etc/match.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	consumer.InitConsumer(ctx)
	logic.InitHandler(ctx)
	logx.Info("server start")
	select {}
}
