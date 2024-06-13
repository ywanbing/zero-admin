// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/productservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type ProductServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductServiceServer
}

func NewProductServiceServer(svcCtx *svc.ServiceContext) *ProductServiceServer {
	return &ProductServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServiceServer) AddProduct(ctx context.Context, in *pmsclient.AddProductReq) (*pmsclient.AddProductResp, error) {
	l := productservicelogic.NewAddProductLogic(ctx, s.svcCtx)
	return l.AddProduct(in)
}

// 查询商品列表
func (s *ProductServiceServer) QueryProductList(ctx context.Context, in *pmsclient.QueryProductListReq) (*pmsclient.QueryProductListResp, error) {
	l := productservicelogic.NewQueryProductListLogic(ctx, s.svcCtx)
	return l.QueryProductList(in)
}

func (s *ProductServiceServer) QueryProductListByIds(ctx context.Context, in *pmsclient.QueryProductByIdsReq) (*pmsclient.QueryProductListResp, error) {
	l := productservicelogic.NewQueryProductListByIdsLogic(ctx, s.svcCtx)
	return l.QueryProductListByIds(in)
}

func (s *ProductServiceServer) UpdateProduct(ctx context.Context, in *pmsclient.UpdateProductReq) (*pmsclient.UpdateProductResp, error) {
	l := productservicelogic.NewUpdateProductLogic(ctx, s.svcCtx)
	return l.UpdateProduct(in)
}

func (s *ProductServiceServer) DeleteProduct(ctx context.Context, in *pmsclient.DeleteProductReq) (*pmsclient.DeleteProductResp, error) {
	l := productservicelogic.NewDeleteProductLogic(ctx, s.svcCtx)
	return l.DeleteProduct(in)
}

// 查询商品详情
func (s *ProductServiceServer) QueryProductDetailById(ctx context.Context, in *pmsclient.QueryProductDetailByIdReq) (*pmsclient.QueryProductDetailByIdResp, error) {
	l := productservicelogic.NewQueryProductDetailByIdLogic(ctx, s.svcCtx)
	return l.QueryProductDetailById(in)
}

// 批量修改审核状态
func (s *ProductServiceServer) UpdateVerifyStatus(ctx context.Context, in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	l := productservicelogic.NewUpdateVerifyStatusLogic(ctx, s.svcCtx)
	return l.UpdateVerifyStatus(in)
}

// 批量上下架商品
func (s *ProductServiceServer) UpdatePublishStatus(ctx context.Context, in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	l := productservicelogic.NewUpdatePublishStatusLogic(ctx, s.svcCtx)
	return l.UpdatePublishStatus(in)
}

// 批量推荐商品
func (s *ProductServiceServer) UpdateRecommendStatus(ctx context.Context, in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	l := productservicelogic.NewUpdateRecommendStatusLogic(ctx, s.svcCtx)
	return l.UpdateRecommendStatus(in)
}

// 批量设为新品
func (s *ProductServiceServer) UpdateNewStatus(ctx context.Context, in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	l := productservicelogic.NewUpdateNewStatusLogic(ctx, s.svcCtx)
	return l.UpdateNewStatus(in)
}

// 批量修改删除状态
func (s *ProductServiceServer) UpdateDeleteStatus(ctx context.Context, in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	l := productservicelogic.NewUpdateDeleteStatusLogic(ctx, s.svcCtx)
	return l.UpdateDeleteStatus(in)
}
