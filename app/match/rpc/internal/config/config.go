package config

import (
	commongorm "gex/common/pkg/gorm"
	"gex/common/pkg/logger"
	"gex/common/pkg/pulsar"
	"gex/common/proto/define"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PulsarConfig pulsar.PulsarConfig
	LoggerConfig logger.Config
	GormConf     commongorm.GormConf
	WsConf       zrpc.RpcClientConf
	SymbolInfo   define.SymbolInfo
	OrderRpcConf zrpc.RpcClientConf
	RedisConf    redis.RedisConf
}
