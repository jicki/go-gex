package main

import (
	"flag"
	"gex/app/order/rpc/internal/config"
	"gex/app/order/rpc/internal/consumer"
	"gex/app/order/rpc/internal/server"
	"gex/app/order/rpc/internal/svc"
	"gex/app/order/rpc/pb"
	"gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/order/rpc/etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(&c)
	consumer.InitConsumer(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServiceServer(grpcServer, server.NewOrderServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.Infof("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
