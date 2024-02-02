package svc

import (
	"gex/app/quotes/kline/rpc/internal/config"
	"gex/app/quotes/kline/rpc/internal/dao/query"
	"gex/common/pkg/logger"
	pulsarConfig "gex/common/pkg/pulsar"
	"github.com/apache/pulsar-client-go/pulsar"
	gpushPb "github.com/luxun9527/gpush/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        *config.Config
	Query         *query.Query
	RedisClient   *redis.Redis
	MatchConsumer pulsar.Consumer
	WsClient      gpushPb.ProxyClient
}

func NewServiceContext(c *config.Config) *ServiceContext {
	logger.InitLogger(c.LoggerConfig)
	logx.SetWriter(logger.NewZapWriter(logger.L))
	logx.DisableStat()
	c.Etcd.Key += "." + c.SymbolInfo.SymbolName
	client, err := c.PulsarConfig.BuildClient()
	if err != nil {
		logx.Severef("init pulsar client failed", logger.ErrorField(err))
	}
	topic := pulsarConfig.Topic{
		Tenant:    pulsarConfig.PublicTenant,
		Namespace: pulsarConfig.GexNamespace,
		Topic:     pulsarConfig.MatchResultTopic + "_" + c.SymbolInfo.SymbolName,
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic.BuildTopic(),
		SubscriptionName: pulsarConfig.MatchResultKlineSub,
		Type:             pulsar.Shared,
	})
	if err != nil {
		logx.Severef("init pulsar consumer failed", logger.ErrorField(err))
	}
	sc := &ServiceContext{
		Config:        c,
		Query:         query.Use(c.GormConf.MustNewGormClient()),
		RedisClient:   redis.MustNewRedis(c.RedisConf),
		MatchConsumer: consumer,
		WsClient:      gpushPb.NewProxyClient(zrpc.MustNewClient(c.WsConf).Conn()),
	}
	return sc
}
