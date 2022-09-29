package bookUse

import (
	"context"

	"go-framework/service/book/api/internal/svc"
	"go-framework/service/book/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BackLogic {
	return &BackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BackLogic) Back(req *types.RentReq) error {
	// todo: add your logic here and delete this line

	return nil
}
