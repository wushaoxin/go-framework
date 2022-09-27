package response

import (
    "github.com/zeromicro/go-zero/rest/httpx"
    "net/http"
)

type Body struct {
    Success bool   `json:"success"`
    Code    int    `json:"code"`
    Msg     string `json:"msg"`
    Data    any    `json:"data,omitempty"`
}

func newSuccessBody(data any) *Body {
    return &Body{
        Success: true,
        Code:    200,
        Msg:     "OK",
        Data:    data,
    }
}

func Response(w http.ResponseWriter, resp any, err error) {
    var body *Body
    if err != nil {
        switch e := err.(type) {
        case *CodeError:
            body = e.Body()
        default:
            body = NewServerError().Body()
        }
        
    } else {
        body = newSuccessBody(resp)
    }
    
    httpx.OkJson(w, body)
}
