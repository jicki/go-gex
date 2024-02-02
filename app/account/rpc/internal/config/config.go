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
	GormConf       commongorm.GormConf
	LoggerConfig   logger.Config
	PulsarConfig   pulsar.PulsarConfig
	RedisConf      redis.RedisConf
	SymbolListConf define.SymbolCoinConfig[string, *define.SymbolInfo]
	CoinListConf   define.SymbolCoinConfig[string, *define.CoinInfo]
}
