// Code generated by goctl. DO NOT EDIT.
// Source: kline.proto

package klineservice

import (
	"context"

	"gex/app/quotes/kline/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetKlineReq        = pb.GetKlineReq
	GetKlineResp       = pb.GetKlineResp
	GetKlineResp_Kline = pb.GetKlineResp_Kline

	KlineService interface {
		// 获取k线
		GetKline(ctx context.Context, in *GetKlineReq, opts ...grpc.CallOption) (*GetKlineResp, error)
	}

	defaultKlineService struct {
		cli zrpc.Client
	}
)

func NewKlineService(cli zrpc.Client) KlineService {
	return &defaultKlineService{
		cli: cli,
	}
}

// 获取k线
func (m *defaultKlineService) GetKline(ctx context.Context, in *GetKlineReq, opts ...grpc.CallOption) (*GetKlineResp, error) {
	client := pb.NewKlineServiceClient(m.cli.Conn())
	return client.GetKline(ctx, in, opts...)
}
