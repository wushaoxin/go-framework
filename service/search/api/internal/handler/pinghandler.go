package handler

import (
	"net/http"

	"go-framework/common/response"
	"go-framework/service/search/api/internal/logic"
	"go-framework/service/search/api/internal/svc"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		response.Response(w, nil, err)
	}
}
