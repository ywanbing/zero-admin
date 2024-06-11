// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/logic/memberreceiveaddressservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
)

type MemberReceiveAddressServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberReceiveAddressServiceServer
}

func NewMemberReceiveAddressServiceServer(svcCtx *svc.ServiceContext) *MemberReceiveAddressServiceServer {
	return &MemberReceiveAddressServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加会员收货地址表
func (s *MemberReceiveAddressServiceServer) AddMemberReceiveAddress(ctx context.Context, in *umsclient.AddMemberReceiveAddressReq) (*umsclient.AddMemberReceiveAddressResp, error) {
	l := memberreceiveaddressservicelogic.NewAddMemberReceiveAddressLogic(ctx, s.svcCtx)
	return l.AddMemberReceiveAddress(in)
}

// 删除会员收货地址表
func (s *MemberReceiveAddressServiceServer) DeleteMemberReceiveAddress(ctx context.Context, in *umsclient.DeleteMemberReceiveAddressReq) (*umsclient.DeleteMemberReceiveAddressResp, error) {
	l := memberreceiveaddressservicelogic.NewDeleteMemberReceiveAddressLogic(ctx, s.svcCtx)
	return l.DeleteMemberReceiveAddress(in)
}

// 更新会员收货地址表
func (s *MemberReceiveAddressServiceServer) UpdateMemberReceiveAddress(ctx context.Context, in *umsclient.UpdateMemberReceiveAddressReq) (*umsclient.UpdateMemberReceiveAddressResp, error) {
	l := memberreceiveaddressservicelogic.NewUpdateMemberReceiveAddressLogic(ctx, s.svcCtx)
	return l.UpdateMemberReceiveAddress(in)
}

// 更新会员收货地址表状态
func (s *MemberReceiveAddressServiceServer) UpdateMemberReceiveAddressStatus(ctx context.Context, in *umsclient.UpdateMemberReceiveAddressStatusReq) (*umsclient.UpdateMemberReceiveAddressStatusResp, error) {
	l := memberreceiveaddressservicelogic.NewUpdateMemberReceiveAddressStatusLogic(ctx, s.svcCtx)
	return l.UpdateMemberReceiveAddressStatus(in)
}

// 查询会员收货地址表详情
func (s *MemberReceiveAddressServiceServer) QueryMemberReceiveAddressDetail(ctx context.Context, in *umsclient.QueryMemberReceiveAddressDetailReq) (*umsclient.QueryMemberReceiveAddressDetailResp, error) {
	l := memberreceiveaddressservicelogic.NewQueryMemberReceiveAddressDetailLogic(ctx, s.svcCtx)
	return l.QueryMemberReceiveAddressDetail(in)
}

// 查询会员收货地址表列表
func (s *MemberReceiveAddressServiceServer) QueryMemberReceiveAddressList(ctx context.Context, in *umsclient.QueryMemberReceiveAddressListReq) (*umsclient.QueryMemberReceiveAddressListResp, error) {
	l := memberreceiveaddressservicelogic.NewQueryMemberReceiveAddressListLogic(ctx, s.svcCtx)
	return l.QueryMemberReceiveAddressList(in)
}
