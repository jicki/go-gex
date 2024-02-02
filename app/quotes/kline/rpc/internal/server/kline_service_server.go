// Code generated by goctl. DO NOT EDIT.
// Source: kline.proto

package server

import (
	"context"

	"gex/app/quotes/kline/rpc/internal/logic"
	"gex/app/quotes/kline/rpc/internal/svc"
	"gex/app/quotes/kline/rpc/pb"
)

type KlineServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedKlineServiceServer
}

func NewKlineServiceServer(svcCtx *svc.ServiceContext) *KlineServiceServer {
	return &KlineServiceServer{
		svcCtx: svcCtx,
	}
}

// 获取k线
func (s *KlineServiceServer) GetKline(ctx context.Context, in *pb.GetKlineReq) (*pb.GetKlineResp, error) {
	l := logic.NewGetKlineLogic(ctx, s.svcCtx)
	return l.GetKline(in)
}
