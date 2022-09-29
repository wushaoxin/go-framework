package main

import (
    "flag"
    "fmt"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/rest/httpx"
    "go-framework/common/response"
    "net/http"
    
    "go-framework/service/book/api/internal/config"
    "go-framework/service/book/api/internal/handler"
    "go-framework/service/book/api/internal/svc"
    
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("book-api-conf", "service/book/api/etc/book-api.yaml", "the config file")

func main() {
    flag.Parse()
    
    var c config.Config
    conf.MustLoad(*configFile, &c)
    
    loadLogConf(&c)
    confGlobalResponse()
    
    server := rest.MustNewServer(c.RestConf)
    defer server.Stop()
    
    ctx := svc.NewServiceContext(c)
    handler.RegisterHandlers(server, ctx)
    
    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}

// 加载日志配置
func loadLogConf(c *config.Config) {
    // 从 yaml 文件中 初始化配置
    err := conf.Load("config.yaml", &c)
    if err != nil {
        fmt.Println("未找到相关日志配置!")
        return
    }
    // logx 根据配置初始化
    logx.MustSetup(c.ServiceConf.Log)
}

func confGlobalResponse() {
    // 自定义错误
    httpx.SetErrorHandler(func(err error) (int, any) {
        switch e := err.(type) {
        case *response.CodeError:
            return http.StatusOK, e.Body()
        default:
            return http.StatusInternalServerError, response.NewServerError().Body()
        }
    })
}
