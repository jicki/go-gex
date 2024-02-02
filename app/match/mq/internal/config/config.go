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
	PulsarConfig pulsar.PulsarConfig
	LoggerConfig logger.Config
	SymbolInfo   define.SymbolInfo
	GormConf     commongorm.GormConf
	WsConf       zrpc.RpcClientConf
	RedisConf    redis.RedisConf
}
