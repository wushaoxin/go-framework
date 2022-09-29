package bookUse

import (
	"context"

	"go-framework/service/book/api/internal/svc"
	"go-framework/service/book/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RentLogic {
	return &RentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RentLogic) Rent(req *types.RentReq) error {
	// todo: add your logic here and delete this line

	return nil
}
