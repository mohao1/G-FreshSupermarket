package LowProduct

import (
	"net/http"

	"DP/DP/internal/logic/LowProduct"
	"DP/DP/internal/svc"
	"DP/DP/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpDateLowProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpDateLowProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := LowProduct.NewUpDateLowProductLogic(r.Context(), svcCtx)
		resp, err := l.UpDateLowProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
