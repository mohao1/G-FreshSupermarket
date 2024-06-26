package Employee

import (
	"net/http"

	"DP/DP/internal/logic/Employee"
	"DP/DP/internal/svc"
	"DP/DP/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteEmployeeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteEmployeeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := Employee.NewDeleteEmployeeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteEmployee(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
