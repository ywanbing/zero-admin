// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/feighttemplateservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type FeightTemplateServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedFeightTemplateServiceServer
}

func NewFeightTemplateServiceServer(svcCtx *svc.ServiceContext) *FeightTemplateServiceServer {
	return &FeightTemplateServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加运费模版
func (s *FeightTemplateServiceServer) AddFeightTemplate(ctx context.Context, in *pmsclient.AddFeightTemplateReq) (*pmsclient.AddFeightTemplateResp, error) {
	l := feighttemplateservicelogic.NewAddFeightTemplateLogic(ctx, s.svcCtx)
	return l.AddFeightTemplate(in)
}

// 删除运费模版
func (s *FeightTemplateServiceServer) DeleteFeightTemplate(ctx context.Context, in *pmsclient.DeleteFeightTemplateReq) (*pmsclient.DeleteFeightTemplateResp, error) {
	l := feighttemplateservicelogic.NewDeleteFeightTemplateLogic(ctx, s.svcCtx)
	return l.DeleteFeightTemplate(in)
}

// 更新运费模版
func (s *FeightTemplateServiceServer) UpdateFeightTemplate(ctx context.Context, in *pmsclient.UpdateFeightTemplateReq) (*pmsclient.UpdateFeightTemplateResp, error) {
	l := feighttemplateservicelogic.NewUpdateFeightTemplateLogic(ctx, s.svcCtx)
	return l.UpdateFeightTemplate(in)
}

// 查询运费模版详情
func (s *FeightTemplateServiceServer) QueryFeightTemplateDetail(ctx context.Context, in *pmsclient.QueryFeightTemplateDetailReq) (*pmsclient.QueryFeightTemplateDetailResp, error) {
	l := feighttemplateservicelogic.NewQueryFeightTemplateDetailLogic(ctx, s.svcCtx)
	return l.QueryFeightTemplateDetail(in)
}

// 查询运费模版列表
func (s *FeightTemplateServiceServer) QueryFeightTemplateList(ctx context.Context, in *pmsclient.QueryFeightTemplateListReq) (*pmsclient.QueryFeightTemplateListResp, error) {
	l := feighttemplateservicelogic.NewQueryFeightTemplateListLogic(ctx, s.svcCtx)
	return l.QueryFeightTemplateList(in)
}
