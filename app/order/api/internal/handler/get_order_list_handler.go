package handler

import (
	"gex/app/order/api/internal/logic"
	"gex/app/order/api/internal/svc"
	"gex/app/order/api/internal/types"
	"gex/common/errs"
	"gex/common/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetOrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, r, nil, errs.WarpMessage(errs.ParamValidateFailed, err.Error()))
			return
		}

		l := logic.NewGetOrderListLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderList(&req)
		response.Response(w, r, resp, err)

	}
}
