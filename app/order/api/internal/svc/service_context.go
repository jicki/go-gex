package svc

import (
	"encoding/json"
	"gex/app/account/rpc/accountservice"
	matchpb "gex/app/match/rpc/pb"
	"gex/app/order/api/internal/config"
	"gex/app/order/api/internal/middleware"
	orderpb "gex/app/order/rpc/pb"
	"gex/common/errs"
	"gex/common/pkg/logger"
	"gex/common/pkg/pool"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"sync"
)

type GetOrderClientFunc func(cc grpc.ClientConnInterface) orderpb.OrderServiceClient
type GetMatchClientFunc func(cc grpc.ClientConnInterface) matchpb.MatchServiceClient

type ServiceContext struct {
	Config           config.Config
	OrderClients     *pool.RpcClients
	MatchClients     *pool.RpcClients
	GetOrderClient   GetOrderClientFunc
	GetMatchClient   GetMatchClientFunc
	Auth             rest.Middleware
	AccountRpcClient accountservice.AccountService
	Symbols          *sync.Map
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()
	d, _ := json.Marshal(c.LanguageEtcdConf)
	errs.InitTranslatorFromEtcd(string(d))

	accountRpcClient := accountservice.NewAccountService(zrpc.MustNewClient(c.AccountRpcConf))
	return &ServiceContext{
		Config:           c,
		Auth:             middleware.NewAuthMiddleware(accountRpcClient).Handle,
		OrderClients:     pool.NewRpcClients(c.OrderRpcConf.Etcd.Key, c.OrderRpcConf.Etcd.Hosts),
		MatchClients:     pool.NewRpcClients(c.MatchRpcConf.Etcd.Key, c.MatchRpcConf.Etcd.Hosts),
		GetOrderClient:   orderpb.NewOrderServiceClient,
		GetMatchClient:   matchpb.NewMatchServiceClient,
		AccountRpcClient: accountRpcClient,
		Symbols:          c.SymbolListConf.CastToSyncMap(),
	}
}
