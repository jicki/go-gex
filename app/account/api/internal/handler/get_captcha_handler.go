package handler

import (
	"gex/app/account/api/internal/logic"
	"gex/app/account/api/internal/svc"
	"gex/common/pkg/response"
	"net/http"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha()
		response.Response(w, r, resp, err) //

	}
}
