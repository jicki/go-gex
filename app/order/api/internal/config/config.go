package config

import (
	"gex/common/pkg/etcd"
	"gex/common/pkg/logger"
	"gex/common/proto/define"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	LoggerConfig     logger.Config
	OrderRpcConf     zrpc.RpcClientConf
	MatchRpcConf     zrpc.RpcClientConf
	AccountRpcConf   zrpc.RpcClientConf
	SymbolListConf   define.SymbolCoinConfig[string, *define.SymbolInfo]
	LanguageEtcdConf etcd.EtcdConfig
}
