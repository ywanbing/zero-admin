package logic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *sys.RoleAddReq) (*sys.RoleAddResp, error) {
	_, _ = l.svcCtx.RoleModel.Insert(l.ctx, &sysmodel.SysRole{
		Id:       0,
		Name:     in.Name,
		Remark:   sql.NullString{String: in.Remark, Valid: true},
		CreateBy: in.CreateBy,
		DelFlag:  0,
		Status:   in.Status,
	})

	return &sys.RoleAddResp{}, nil
}
