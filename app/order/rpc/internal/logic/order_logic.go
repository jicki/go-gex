package logic

import (
	"context"
	accountpb "gex/app/account/rpc/pb"
	"gex/app/order/rpc/internal/svc"
	"gex/app/order/rpc/pb"
	"gex/common/errs"
	"gex/common/pkg/logger"
	enum "gex/common/proto/enum"
	"gex/common/utils"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

type OrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 下单。
func (l *OrderLogic) Order(in *pb.CreateOrderReq) (*pb.OrderEmpty, error) {

	freezeReq := &accountpb.FreezeUserAssetReq{
		Uid:    in.UserId,
		CoinId: 0,
		Qty:    in.Qty,
	}
	//如果是买冻结计价币
	if in.Side == enum.Side_Buy {
		freezeReq.CoinId = l.svcCtx.Config.SymbolInfo.QuoteCoinID
		freezeReq.Qty = in.Amount
		if in.OrderType == enum.OrderType_LO {
			freezeReq.Qty = utils.NewFromStringMaxPrec(in.Qty).Mul(utils.NewFromStringMaxPrec(in.Price)).String()
		}
	} else {
		freezeReq.CoinId = l.svcCtx.Config.SymbolInfo.BaseCoinID
	}
	createOrderReq := &pb.CreateOrderReq{
		UserId:     in.UserId,
		SymbolId:   in.SymbolId,
		SymbolName: in.SymbolName,
		Qty:        in.Qty,
		Price:      in.Price,
		Side:       in.Side,
		OrderType:  in.OrderType,
		Amount:     in.Amount,
	}
	gid, err := l.svcCtx.DtmClient.NewGid(l.ctx, &emptypb.Empty{})
	if err != nil {
		logx.Errorw("get gid failed", logger.ErrorField(err))
		return nil, errs.DTMFailed
	}
	//l.svcCtx.Config.AccountRpcConf.
	//gid := stringx.Randn(16)
	//	l.svcCtx.Config.AccountRpcConf.Target
	accountTarget, err := l.svcCtx.Config.AccountRpcConf.BuildTarget()
	if err != nil {
		logx.Errorw("get account client failed", logger.ErrorField(err))
		return nil, errs.Internal
	}
	orderTarget, err := l.svcCtx.Config.OrderRpcConf.BuildTarget()
	if err != nil {
		logx.Errorw("get order client failed", logger.ErrorField(err))
		return nil, errs.Internal
	}
	//l.svcCtx.Config.OrderRpcConf.Endpoints
	var (
		freezeUserAddr    = accountTarget + "/account.AccountService/FreezeUserAsset"
		unFreezeUserAddr  = accountTarget + "/account.AccountService/UnFreezeUserAsset"
		createOrderAddr   = orderTarget + "/order.OrderService/CreateOrder"
		createOrderRevert = orderTarget + "/order.OrderService/CreateOrderRevert"
	)

	dtmAddr, err := l.svcCtx.Config.DtmConf.BuildTarget()
	if err != nil {
		return nil, err
	}
	sagaGrpc := dtmgrpc.NewSagaGrpc(dtmAddr, gid.Gid)
	sagaGrpc.WaitResult = true
	if err := sagaGrpc.
		Add(freezeUserAddr, unFreezeUserAddr, freezeReq).
		Add(createOrderAddr, createOrderRevert, createOrderReq).Submit(); err != nil {
		logx.Errorw("Submit saga  failed", logger.ErrorField(err))
		s, ok := status.FromError(err)
		if ok {
			if s.Code() == codes.Aborted && strings.Contains(s.Message(), "=") {
				msg := s.Message()
				start, end := strings.Index(msg, "="), strings.LastIndex(msg, "=")
				d := msg[start+1 : end]
				e, err := cast.ToInt32E(d)
				if err != nil {
					return nil, errs.Internal
				}
				return nil, errs.Code(e).NewError()
			}

		}
		return nil, errs.DTMFailed
	}
	return &pb.OrderEmpty{}, nil
}
