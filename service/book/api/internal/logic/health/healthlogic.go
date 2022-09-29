package health

import (
    "context"
    
    "github.com/zeromicro/go-zero/core/logx"
    "go-framework/service/book/api/internal/svc"
)

type HealthLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
    return &HealthLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *HealthLogic) Health() error {
    panic("测试错误日志")
    return nil
}
