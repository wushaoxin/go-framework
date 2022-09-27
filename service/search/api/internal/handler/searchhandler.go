package handler

import (
    "go-framework/common/response"
    
    "net/http"
    
    "github.com/zeromicro/go-zero/rest/httpx"
    "go-framework/service/search/api/internal/logic"
    "go-framework/service/search/api/internal/svc"
    "go-framework/service/search/api/internal/types"
)

func searchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.SearchReq
        if err := httpx.Parse(r, &req); err != nil {
            response.Response(w, nil, err)
            return
        }
        
        l := logic.NewSearchLogic(r.Context(), svcCtx)
        resp, err := l.Search(&req)
        response.Response(w, resp, err)
    }
}
