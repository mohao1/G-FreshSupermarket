package StoreServices

import (
	"net/http"

	"DP/DP/internal/logic/StoreServices"
	"DP/DP/internal/svc"
	"DP/DP/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAnnouncementListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnnouncementListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := StoreServices.NewGetAnnouncementListLogic(r.Context(), svcCtx)
		resp, err := l.GetAnnouncementList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
