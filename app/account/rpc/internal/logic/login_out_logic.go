package logic

import (
	"context"
	"gex/common/errs"
	"gex/common/pkg/logger"
	"gex/common/utils"
	"github.com/spf13/cast"
	"time"

	"gex/app/account/rpc/internal/svc"
	"gex/app/account/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginOutLogic {
	return &LoginOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登出 使用黑名单的策略
func (l *LoginOutLogic) LoginOut(in *pb.LoginOutReq) (*pb.Empty, error) {
	key := utils.GenerateKey(in.Token)
	t := cast.ToString(time.Now().Unix())
	token, err := l.svcCtx.JwtClient.ParseToken(in.Token)
	if err != nil {
		return nil, errs.TokenValidateFailed
	}
	remain := token.ExpiresAt.Sub(time.Now()) / 1e9
	if _, err := l.svcCtx.RedisClient.SetnxEx(key, t, int(remain)); err != nil {
		logx.Errorw("set redis failed", logger.ErrorField(err))
		return nil, errs.RedisFailed
	}
	return &pb.Empty{}, nil
}
