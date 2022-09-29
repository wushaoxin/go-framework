package svc

import (
    "github.com/zeromicro/go-zero/rest"
    "github.com/zeromicro/go-zero/zrpc"
    "go-framework/service/book/api/internal/config"
    "go-framework/service/book/api/internal/middleware"
    "go-framework/service/user/rpc/userclient"
)

type ServiceContext struct {
    Config  config.Config
    Example rest.Middleware
    UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:  c,
        Example: middleware.NewExampleMiddleware().Handle,
        UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
    }
}
