package consumer

import (
	"context"
	"gex/app/match/rpc/internal/engine"
	"gex/app/match/rpc/internal/svc"
	"gex/common/pkg/logger"
	"gex/common/proto/enum"
	matchMq "gex/common/proto/mq/match"
	"gex/common/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

func InitMatchConsumer(sc *svc.ServiceContext) {
	ctx := context.Background()
	go func() {
		for {
			message, err := sc.MatchConsumer.Receive(ctx)
			if err != nil {
				logx.Errorw("receive message fail", logger.ErrorField(err))
				continue
			}
			var matchReq matchMq.MatchReq
			if err := proto.Unmarshal(message.Payload(), &matchReq); err != nil {
				logx.Errorw("unmarshal message fail", logger.ErrorField(err))
				continue
			}
			logx.Infow("receive message failed", logx.Field("data", matchReq))
			switch operate := matchReq.Operate.(type) {
			case *matchMq.MatchReq_NewOrder:
				order := &engine.Order{
					Uid:            operate.NewOrder.Uid,
					OrderID:        operate.NewOrder.OrderId,
					SequenceId:     operate.NewOrder.SequenceId,
					CreateTime:     0,
					IsCancel:       false,
					Price:          utils.NewFromStringMaxPrec(operate.NewOrder.Price),
					Qty:            utils.NewFromStringMaxPrec(operate.NewOrder.Qty),
					OrderType:      operate.NewOrder.OrderType,
					Amount:         utils.NewFromStringMaxPrec(operate.NewOrder.Amount),
					Side:           operate.NewOrder.Side,
					OrderStatus:    enum.OrderStatus_NewCreated,
					UnfilledQty:    utils.NewFromStringMaxPrec(operate.NewOrder.Qty),
					FilledQty:      utils.DecimalZeroMaxPrec,
					UnfilledAmount: utils.NewFromStringMaxPrec(operate.NewOrder.Amount),
					FilledAmount:   utils.DecimalZeroMaxPrec,
				}
				sc.MatchEngine.HandleOrder(order)
			case *matchMq.MatchReq_Cancel:
				order := &engine.Order{
					OrderID:    "",
					SequenceId: operate.Cancel.Id,
					CreateTime: 0,
					IsCancel:   true,
					Side:       operate.Cancel.Side,
					Uid:        0,
					OrderType:  operate.Cancel.OrderType,
					Price:      utils.NewFromStringMaxPrec(operate.Cancel.Price),
				}
				sc.MatchEngine.HandleOrder(order)
			}
			if err := sc.MatchConsumer.Ack(message); err != nil {
				logx.Errorw("consumer message failed", logger.ErrorField(err))
			}
		}
	}()
}
