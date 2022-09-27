package main

import (
    "flag"
    "fmt"
    
    "go-framework/service/user/rpc/internal/config"
    "go-framework/service/user/rpc/internal/server"
    "go-framework/service/user/rpc/internal/svc"
    "go-framework/service/user/rpc/user"
    
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/core/service"
    "github.com/zeromicro/go-zero/zrpc"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/user/rpc/etc/user-rpc.yaml", "the config file")

func main() {
    flag.Parse()
    
    var c config.Config
    conf.MustLoad(*configFile, &c)
    ctx := svc.NewServiceContext(c)
    
    s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
        user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))
        
        if c.Mode == service.DevMode || c.Mode == service.TestMode {
            reflection.Register(grpcServer)
        }
    })
    defer s.Stop()
    
    fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
    s.Start()
}
