// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/logic/flashpromotionproductrelationservice"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
)

type FlashPromotionProductRelationServiceServer struct {
	svcCtx *svc.ServiceContext
	smsclient.UnimplementedFlashPromotionProductRelationServiceServer
}

func NewFlashPromotionProductRelationServiceServer(svcCtx *svc.ServiceContext) *FlashPromotionProductRelationServiceServer {
	return &FlashPromotionProductRelationServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加商品限时购与商品关系表
func (s *FlashPromotionProductRelationServiceServer) AddFlashPromotionProductRelation(ctx context.Context, in *smsclient.AddFlashPromotionProductRelationReq) (*smsclient.AddFlashPromotionProductRelationResp, error) {
	l := flashpromotionproductrelationservicelogic.NewAddFlashPromotionProductRelationLogic(ctx, s.svcCtx)
	return l.AddFlashPromotionProductRelation(in)
}

// 删除商品限时购与商品关系表
func (s *FlashPromotionProductRelationServiceServer) DeleteFlashPromotionProductRelation(ctx context.Context, in *smsclient.DeleteFlashPromotionProductRelationReq) (*smsclient.DeleteFlashPromotionProductRelationResp, error) {
	l := flashpromotionproductrelationservicelogic.NewDeleteFlashPromotionProductRelationLogic(ctx, s.svcCtx)
	return l.DeleteFlashPromotionProductRelation(in)
}

// 更新商品限时购与商品关系表
func (s *FlashPromotionProductRelationServiceServer) UpdateFlashPromotionProductRelation(ctx context.Context, in *smsclient.UpdateFlashPromotionProductRelationReq) (*smsclient.UpdateFlashPromotionProductRelationResp, error) {
	l := flashpromotionproductrelationservicelogic.NewUpdateFlashPromotionProductRelationLogic(ctx, s.svcCtx)
	return l.UpdateFlashPromotionProductRelation(in)
}

// 查询商品限时购与商品关系表详情
func (s *FlashPromotionProductRelationServiceServer) QueryFlashPromotionProductRelationDetail(ctx context.Context, in *smsclient.QueryFlashPromotionProductRelationDetailReq) (*smsclient.QueryFlashPromotionProductRelationDetailResp, error) {
	l := flashpromotionproductrelationservicelogic.NewQueryFlashPromotionProductRelationDetailLogic(ctx, s.svcCtx)
	return l.QueryFlashPromotionProductRelationDetail(in)
}

// 查询商品限时购与商品关系表列表
func (s *FlashPromotionProductRelationServiceServer) QueryFlashPromotionProductRelationList(ctx context.Context, in *smsclient.QueryFlashPromotionProductRelationListReq) (*smsclient.QueryFlashPromotionProductRelationListResp, error) {
	l := flashpromotionproductrelationservicelogic.NewQueryFlashPromotionProductRelationListLogic(ctx, s.svcCtx)
	return l.QueryFlashPromotionProductRelationList(in)
}
