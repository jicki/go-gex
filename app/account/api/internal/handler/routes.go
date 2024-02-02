// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/luxun9527/gex/app/account/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/get_captcha",
				Handler: GetCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/validate_token",
				Handler: ValidateTokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/account/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/get_user_asset_list",
					Handler: GetUserAssetListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add_user_asset",
					Handler: AddUserAssetHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/account/v1"),
	)
}
