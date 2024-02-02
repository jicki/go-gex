package handler

import (
	"gex/app/admin/api/internal/logic"
	"gex/app/admin/api/internal/svc"
	"gex/app/admin/api/internal/types"
	"gex/common/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UpsertServiceConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpsertServiceConfigReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, err)
			return
		}

		l := logic.NewUpsertServiceConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpsertServiceConfig(&req)
		response.Response(w, r, resp, err) //

	}
}
