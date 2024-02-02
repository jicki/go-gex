// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/luxun9527/gex/app/quotes/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/get_kline_list",
				Handler: GetKlineListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/get_depth_list",
				Handler: GetDepthListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/get_ticker_list",
				Handler: GetTickerListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/get_tick_list",
				Handler: GetTickListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/quotes/v1"),
	)
}