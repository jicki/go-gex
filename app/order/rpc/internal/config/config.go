package config

import (
	commongorm "gex/common/pkg/gorm"
	"gex/common/pkg/logger"
	"gex/common/pkg/pulsar"
	"gex/common/proto/define"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AccountRpcConf  zrpc.RpcClientConf
	OrderRpcConf    zrpc.RpcClientConf
	DtmConf         zrpc.RpcClientConf
	PulsarConfig    pulsar.PulsarConfig
	LoggerConfig    logger.Config
	GormConf        commongorm.GormConf
	SymbolInfo      define.SymbolInfo
	SnowFlakeWorkID int64
	WsConf          zrpc.RpcClientConf
}
