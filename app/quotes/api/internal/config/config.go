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
	KlineRpcConf     zrpc.RpcClientConf
	MatchRpcConf     zrpc.RpcClientConf
	SymbolList       []*define.SymbolInfo
	LoggerConfig     logger.Config
	LanguageEtcdConf etcd.EtcdConfig
}
