package logic

import (
	"context"
	"github.com/luxun9527/gex/app/account/rpc/accountservice"
	"github.com/luxun9527/gex/common/errs"
	"github.com/spf13/cast"

	"github.com/luxun9527/gex/app/account/api/internal/svc"
	"github.com/luxun9527/gex/app/account/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidateTokenLogic) ValidateToken(req *types.ValidateTokenReq) (resp *types.ValidateTokenResp, err error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.AccountRpcClient.ValidateToken(l.ctx, &accountservice.ValidateTokenReq{Token: req.Token})
	if err != nil {
		return nil, errs.TokenValidateFailed
	}
	resp = &types.ValidateTokenResp{UserInfo: &types.UserInfo{
		Uid:      cast.ToString(userInfo.Uid),
		Username: userInfo.Username,
	}}
	return
}
