package config

import (
	"gex/common/pkg/etcd"
	"gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	LoggerConfig     logger.Config
	AccountRpcConf   zrpc.RpcClientConf
	LanguageEtcdConf etcd.EtcdConfig
}
