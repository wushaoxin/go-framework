package api

import (
    "flag"
    "fmt"
    "github.com/zeromicro/go-zero/core/logx"
    "net/http"
    
    "go-framework/service/search/api/internal/config"
    "go-framework/service/search/api/internal/handler"
    "go-framework/service/search/api/internal/svc"
    
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("search-api-conf", "service/search/api/etc/search-api.yaml", "the config file")

func Main() {
    flag.Parse()
    
    var c config.Config
    conf.MustLoad(*configFile, &c)
    
    ctx := svc.NewServiceContext(c)
    server := rest.MustNewServer(c.RestConf, rest.WithCors())
    defer server.Stop()
    
    // 全局中间件
    server.Use(func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            logx.Info("global middleware")
            next(w, r)
        }
    })
    handler.RegisterHandlers(server, ctx)
    
    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
