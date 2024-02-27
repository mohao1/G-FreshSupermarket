package AdminProductType

import (
	"net/http"

	"DP/DP/internal/logic/AdminProductType"
	"DP/DP/internal/svc"
	"DP/DP/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteProductTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductTypeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := AdminProductType.NewDeleteProductTypeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProductType(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
