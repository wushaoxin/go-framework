package handler

import (
    "go-framework/common/response"
    
    "net/http"
    
    "github.com/zeromicro/go-zero/rest/httpx"
    "go-framework/service/user/api/internal/logic"
    "go-framework/service/user/api/internal/svc"
    "go-framework/service/user/api/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.LoginReq
        if err := httpx.Parse(r, &req); err != nil {
            response.Response(w, nil, err)
            return
        }
        
        l := logic.NewLoginLogic(r.Context(), svcCtx)
        resp, err := l.Login(&req)
        response.Response(w, resp, err)
    }
}
