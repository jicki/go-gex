package logic

import (
	"context"
	"gex/common/errs"
	"gex/common/pkg/logger"
	"gex/common/utils"

	"gex/app/account/rpc/internal/svc"
	"gex/app/account/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证token是否有效。
func (l *ValidateTokenLogic) ValidateToken(in *pb.ValidateTokenReq) (*pb.ValidateTokenResp, error) {
	claims, err := l.svcCtx.JwtClient.ParseToken(in.Token)
	if err != nil {
		logx.Errorw("parse token failed", logger.ErrorField(err), logx.Field("token", in.Token))
		return nil, errs.Internal
	}

	key := utils.GenerateKey(in.Token)
	result, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		logx.Errorw("get redis key failed", logger.ErrorField(err), logx.Field("key", key))
		return nil, errs.RedisFailed
	}
	if result != "" {
		return nil, errs.TokenValidateFailed
	}
	return &pb.ValidateTokenResp{
		Uid:      claims.UserID,
		Username: claims.Username,
	}, nil
}
