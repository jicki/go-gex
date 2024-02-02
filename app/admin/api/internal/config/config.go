package config

import (
	"gex/common/pkg/etcd"
	"gex/common/pkg/logger"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	LoggerConfig logger.Config
	EtcdConf     etcd.EtcdConfig
}
