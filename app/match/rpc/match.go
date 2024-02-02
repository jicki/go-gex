package main

import (
	"flag"
	"gex/app/match/rpc/internal/bootstrap"
	"gex/app/match/rpc/internal/config"
	"gex/app/match/rpc/internal/server"
	"gex/app/match/rpc/internal/svc"
	"gex/app/match/rpc/pb"
	"gex/common/pkg/confx"
	"gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	symbol     = flag.String("s", "BTC_USDT", "symbol 交易对")
	etcdConfig = flag.String("e", `{"Endpoints":["etcd:2379"],"DialTimeout":5}`, "symbol 交易对")
)

func main() {
	flag.Parse()

	var c config.Config
	//初始化配置
	confx.MustLoadFromEtcd(confx.Match.BuildKey(*symbol), *etcdConfig, &c, confx.WithDefaultInitLoadFunc())
	ctx := svc.NewServiceContext(&c)
	//初始化
	bootstrap.Start(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMatchServiceServer(grpcServer, server.NewMatchServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.Infof("Starting rpc server at %s...", c.ListenOn)
	s.Start()
}
