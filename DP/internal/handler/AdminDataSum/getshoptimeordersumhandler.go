package AdminDataSum

import (
	"net/http"

	"DP/DP/internal/logic/AdminDataSum"
	"DP/DP/internal/svc"
	"DP/DP/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetShopTimeOrderSumHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetShopTimeOrderSumReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := AdminDataSum.NewGetShopTimeOrderSumLogic(r.Context(), svcCtx)
		resp, err := l.GetShopTimeOrderSum(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
