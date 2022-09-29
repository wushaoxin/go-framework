package bookManager

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"go-framework/common/response"
	"go-framework/service/book/api/internal/logic/bookManager"
	"go-framework/service/book/api/internal/svc"
	"go-framework/service/book/api/internal/types"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := bookManager.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		response.Response(w, resp, err)
	}
}
