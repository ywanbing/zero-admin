package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberStatisticsInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberStatisticsInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberStatisticsInfoUpdateLogic {
	return MemberStatisticsInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberStatisticsInfoUpdateLogic) MemberStatisticsInfoUpdate(req types.UpdateMemberStatisticsInfoReq) (*types.UpdateMemberStatisticsInfoResp, error) {
	// todo: add your logic here and delete this line

	return &types.UpdateMemberStatisticsInfoResp{}, nil
}
