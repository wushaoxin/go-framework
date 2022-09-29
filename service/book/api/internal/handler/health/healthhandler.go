package health

import (
	"net/http"

	"go-framework/common/response"
	"go-framework/service/book/api/internal/logic/health"
	"go-framework/service/book/api/internal/svc"
)

func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := health.NewHealthLogic(r.Context(), svcCtx)
		err := l.Health()
		response.Response(w, nil, err)
	}
}
