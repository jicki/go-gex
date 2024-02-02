package logic

import (
	"context"
	"gex/app/admin/api/internal/svc"
	"gex/app/admin/api/internal/types"
	"gex/common/errs"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertServiceConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpsertServiceConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertServiceConfigLogic {
	return &UpsertServiceConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpsertServiceConfigLogic) UpsertServiceConfig(req *types.UpsertServiceConfigReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line
	if _, err := l.svcCtx.EtcdCli.Put(l.ctx, req.Key, req.ConfigData); err != nil {
		logx.Errorw("put config to etcd failed", logx.Field("err", err))
		return nil, errs.Internal
	}

	return
}
