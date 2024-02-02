package handler

import (
	"gex/app/account/api/internal/logic"
	"gex/app/account/api/internal/svc"
	"gex/app/account/api/internal/types"
	"gex/common/errs"
	"gex/common/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AddUserAssetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserAssetReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, errs.WarpMessage(errs.ParamValidateFailed, err.Error()))
			return
		}

		l := logic.NewAddUserAssetLogic(r.Context(), svcCtx)
		resp, err := l.AddUserAsset(&req)
		response.Response(w, r, resp, err)

	}
}
