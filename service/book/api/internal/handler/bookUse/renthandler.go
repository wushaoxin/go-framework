package bookUse

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"go-framework/common/response"
	"go-framework/service/book/api/internal/logic/bookUse"
	"go-framework/service/book/api/internal/svc"
	"go-framework/service/book/api/internal/types"
)

func RentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RentReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := bookUse.NewRentLogic(r.Context(), svcCtx)
		err := l.Rent(&req)
		response.Response(w, nil, err)
	}
}
