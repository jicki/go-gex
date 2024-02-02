package main

import (
	"flag"
	"gex/common/pkg/logger"
	"gex/common/pkg/validatorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"gex/app/order/api/internal/config"
	"gex/app/order/api/internal/handler"
	"gex/app/order/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/order/api/etc/order_api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	httpx.SetValidator(validatorx.NewValidator())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
