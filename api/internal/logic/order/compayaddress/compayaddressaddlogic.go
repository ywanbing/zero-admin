package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CompayAddressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressAddLogic {
	return CompayAddressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressAddLogic) CompayAddressAdd(req types.AddCompayAddressReq) (*types.AddCompayAddressResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddCompayAddressResp{}, nil
}