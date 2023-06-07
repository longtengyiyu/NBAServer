package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"nbaserver/internal/logic"
	"nbaserver/internal/svc"
	"nbaserver/internal/types"
)

func NbaserverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewNbaserverLogic(r.Context(), svcCtx)
		resp, err := l.Nbaserver(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
