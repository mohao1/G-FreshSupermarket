package OrderServices

import (
	"net/http"

	"DP/DP/internal/logic/OrderServices"
	"DP/DP/internal/svc"
	"DP/DP/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OverOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OverOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := OrderServices.NewOverOrderLogic(r.Context(), svcCtx)
		resp, err := l.OverOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
