package main

import (
    "flag"
    "fmt"
    "github.com/zeromicro/go-zero/rest/httpx"
    "go-framework/common/response"
    "net/http"
    
    "go-framework/service/user/api/internal/config"
    "go-framework/service/user/api/internal/handler"
    "go-framework/service/user/api/internal/svc"
    
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "service/user/api/etc/user-api.yaml", "the config file")

func main() {
    flag.Parse()
    
    var c config.Config
    conf.MustLoad(*configFile, &c)
    
    server := rest.MustNewServer(c.RestConf, rest.WithCors())
    defer server.Stop()
    
    ctx := svc.NewServiceContext(c)
    handler.RegisterHandlers(server, ctx)
    
    // 自定义错误
    httpx.SetErrorHandler(func(err error) (int, any) {
        switch e := err.(type) {
        case *response.CodeError:
            return http.StatusOK, e.Body()
        default:
            return http.StatusInternalServerError, response.NewServerError().Body()
        }
    })
    
    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
