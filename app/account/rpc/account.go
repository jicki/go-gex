package main

import (
	"flag"
	"gex/app/account/rpc/internal/config"
	"gex/app/account/rpc/internal/consumer"
	"gex/app/account/rpc/internal/server"
	"gex/app/account/rpc/internal/svc"
	"gex/app/account/rpc/pb"
	"gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "app/account/rpc/etc/account.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	consumer.InitConsumer(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAccountServiceServer(grpcServer, server.NewAccountServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.Infof("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
