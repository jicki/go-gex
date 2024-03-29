package handler

import (
	"gex/app/quotes/api/internal/logic"
	"gex/app/quotes/api/internal/svc"
	"gex/app/quotes/api/internal/types"
	"gex/common/errs"
	"gex/common/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetTickerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTickerListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, errs.WarpMessage(errs.ParamValidateFailed, err.Error()))
			return
		}

		l := logic.NewGetTickerListLogic(r.Context(), svcCtx)
		resp, err := l.GetTickerList(&req)
		response.Response(w, r, resp, err)

	}
}
