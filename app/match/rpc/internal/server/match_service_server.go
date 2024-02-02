// Code generated by goctl. DO NOT EDIT.
// Source: match.proto

package server

import (
	"context"

	"gex/app/match/rpc/internal/logic"
	"gex/app/match/rpc/internal/svc"
	"gex/app/match/rpc/pb"
)

type MatchServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedMatchServiceServer
}

func NewMatchServiceServer(svcCtx *svc.ServiceContext) *MatchServiceServer {
	return &MatchServiceServer{
		svcCtx: svcCtx,
	}
}

// 获取深度
func (s *MatchServiceServer) GetDepth(ctx context.Context, in *pb.GetDepthReq) (*pb.GetDepthResp, error) {
	l := logic.NewGetDepthLogic(ctx, s.svcCtx)
	return l.GetDepth(in)
}

// 获取tick实时成交
func (s *MatchServiceServer) GetTick(ctx context.Context, in *pb.GetTickReq) (*pb.GetTickResp, error) {
	l := logic.NewGetTickLogic(ctx, s.svcCtx)
	return l.GetTick(in)
}

// 获取ticker
func (s *MatchServiceServer) GetTicker(ctx context.Context, in *pb.GetTickerReq) (*pb.GetTickerResp, error) {
	l := logic.NewGetTickerLogic(ctx, s.svcCtx)
	return l.GetTicker(in)
}
