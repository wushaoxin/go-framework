package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"go-framework/common/response"
	"go-framework/service/user/api/internal/logic/user"
	"go-framework/service/user/api/internal/svc"
	"go-framework/service/user/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)
	}
}
