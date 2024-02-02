package main

import (
	"flag"
	"gex/app/quotes/api/internal/config"
	"gex/app/quotes/api/internal/handler"
	"gex/app/quotes/api/internal/svc"
	"gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/quotes/api/etc/quote.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.DisableStat()
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
