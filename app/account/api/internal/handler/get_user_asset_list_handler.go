package handler

import (
	"gex/app/account/api/internal/logic"
	"gex/app/account/api/internal/svc"
	"gex/common/pkg/response"
	"net/http"
)

func GetUserAssetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetUserAssetListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserAssetList()
		response.Response(w, r, resp, err) //

	}
}
