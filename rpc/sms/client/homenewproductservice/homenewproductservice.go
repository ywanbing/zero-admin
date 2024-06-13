// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package homenewproductservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCouponHistoryReq                           = smsclient.AddCouponHistoryReq
	AddCouponHistoryResp                          = smsclient.AddCouponHistoryResp
	AddFlashPromotionLogReq                       = smsclient.AddFlashPromotionLogReq
	AddFlashPromotionLogResp                      = smsclient.AddFlashPromotionLogResp
	AddFlashPromotionProductRelationReq           = smsclient.AddFlashPromotionProductRelationReq
	AddFlashPromotionProductRelationResp          = smsclient.AddFlashPromotionProductRelationResp
	AddFlashPromotionReq                          = smsclient.AddFlashPromotionReq
	AddFlashPromotionResp                         = smsclient.AddFlashPromotionResp
	AddFlashPromotionSessionReq                   = smsclient.AddFlashPromotionSessionReq
	AddFlashPromotionSessionResp                  = smsclient.AddFlashPromotionSessionResp
	AddHomeAdvertiseReq                           = smsclient.AddHomeAdvertiseReq
	AddHomeAdvertiseResp                          = smsclient.AddHomeAdvertiseResp
	AddHomeBrandReq                               = smsclient.AddHomeBrandReq
	AddHomeBrandResp                              = smsclient.AddHomeBrandResp
	AddHomeNewProductReq                          = smsclient.AddHomeNewProductReq
	AddHomeNewProductResp                         = smsclient.AddHomeNewProductResp
	AddHomeRecommendProductReq                    = smsclient.AddHomeRecommendProductReq
	AddHomeRecommendProductResp                   = smsclient.AddHomeRecommendProductResp
	AddHomeRecommendSubjectReq                    = smsclient.AddHomeRecommendSubjectReq
	AddHomeRecommendSubjectResp                   = smsclient.AddHomeRecommendSubjectResp
	AddOrUpdateCouponReq                          = smsclient.AddOrUpdateCouponReq
	AddOrUpdateCouponResp                         = smsclient.AddOrUpdateCouponResp
	CouponFindByProductIdAndProductCategoryIdReq  = smsclient.CouponFindByProductIdAndProductCategoryIdReq
	CouponFindByProductIdAndProductCategoryIdResp = smsclient.CouponFindByProductIdAndProductCategoryIdResp
	CouponHistoryDetailListData                   = smsclient.CouponHistoryDetailListData
	CouponHistoryDetailListResp                   = smsclient.CouponHistoryDetailListResp
	CouponHistoryListData                         = smsclient.CouponHistoryListData
	CouponProductCategoryRelationData             = smsclient.CouponProductCategoryRelationData
	CouponProductRelationData                     = smsclient.CouponProductRelationData
	DeleteCouponHistoryReq                        = smsclient.DeleteCouponHistoryReq
	DeleteCouponHistoryResp                       = smsclient.DeleteCouponHistoryResp
	DeleteCouponReq                               = smsclient.DeleteCouponReq
	DeleteCouponResp                              = smsclient.DeleteCouponResp
	DeleteFlashPromotionLogReq                    = smsclient.DeleteFlashPromotionLogReq
	DeleteFlashPromotionLogResp                   = smsclient.DeleteFlashPromotionLogResp
	DeleteFlashPromotionProductRelationReq        = smsclient.DeleteFlashPromotionProductRelationReq
	DeleteFlashPromotionProductRelationResp       = smsclient.DeleteFlashPromotionProductRelationResp
	DeleteFlashPromotionReq                       = smsclient.DeleteFlashPromotionReq
	DeleteFlashPromotionResp                      = smsclient.DeleteFlashPromotionResp
	DeleteFlashPromotionSessionReq                = smsclient.DeleteFlashPromotionSessionReq
	DeleteFlashPromotionSessionResp               = smsclient.DeleteFlashPromotionSessionResp
	DeleteHomeAdvertiseReq                        = smsclient.DeleteHomeAdvertiseReq
	DeleteHomeAdvertiseResp                       = smsclient.DeleteHomeAdvertiseResp
	DeleteHomeBrandReq                            = smsclient.DeleteHomeBrandReq
	DeleteHomeBrandResp                           = smsclient.DeleteHomeBrandResp
	DeleteHomeNewProductReq                       = smsclient.DeleteHomeNewProductReq
	DeleteHomeNewProductResp                      = smsclient.DeleteHomeNewProductResp
	DeleteHomeRecommendProductReq                 = smsclient.DeleteHomeRecommendProductReq
	DeleteHomeRecommendProductResp                = smsclient.DeleteHomeRecommendProductResp
	DeleteHomeRecommendSubjectReq                 = smsclient.DeleteHomeRecommendSubjectReq
	DeleteHomeRecommendSubjectResp                = smsclient.DeleteHomeRecommendSubjectResp
	FlashPromotionListData                        = smsclient.FlashPromotionListData
	FlashPromotionLogListData                     = smsclient.FlashPromotionLogListData
	FlashPromotionProductRelationAddData          = smsclient.FlashPromotionProductRelationAddData
	FlashPromotionProductRelationListData         = smsclient.FlashPromotionProductRelationListData
	FlashPromotionSessionListData                 = smsclient.FlashPromotionSessionListData
	HomeAdvertiseListData                         = smsclient.HomeAdvertiseListData
	HomeBrandAddData                              = smsclient.HomeBrandAddData
	HomeBrandListData                             = smsclient.HomeBrandListData
	HomeNewProductAddData                         = smsclient.HomeNewProductAddData
	HomeNewProductListData                        = smsclient.HomeNewProductListData
	HomeRecommendProductAddData                   = smsclient.HomeRecommendProductAddData
	HomeRecommendProductListData                  = smsclient.HomeRecommendProductListData
	HomeRecommendSubjectAddData                   = smsclient.HomeRecommendSubjectAddData
	HomeRecommendSubjectListData                  = smsclient.HomeRecommendSubjectListData
	QueryCouponCountReq                           = smsclient.QueryCouponCountReq
	QueryCouponCountResp                          = smsclient.QueryCouponCountResp
	QueryCouponData                               = smsclient.QueryCouponData
	QueryCouponFindByIdReq                        = smsclient.QueryCouponFindByIdReq
	QueryCouponFindByIdResp                       = smsclient.QueryCouponFindByIdResp
	QueryCouponFindByIdsReq                       = smsclient.QueryCouponFindByIdsReq
	QueryCouponFindByIdsResp                      = smsclient.QueryCouponFindByIdsResp
	QueryCouponHistoryDetailListReq               = smsclient.QueryCouponHistoryDetailListReq
	QueryCouponHistoryDetailReq                   = smsclient.QueryCouponHistoryDetailReq
	QueryCouponHistoryDetailResp                  = smsclient.QueryCouponHistoryDetailResp
	QueryCouponHistoryListReq                     = smsclient.QueryCouponHistoryListReq
	QueryCouponHistoryListResp                    = smsclient.QueryCouponHistoryListResp
	QueryCouponListData                           = smsclient.QueryCouponListData
	QueryCouponListReq                            = smsclient.QueryCouponListReq
	QueryCouponListResp                           = smsclient.QueryCouponListResp
	QueryCouponProductCategoryRelationListData    = smsclient.QueryCouponProductCategoryRelationListData
	QueryCouponProductRelationListData            = smsclient.QueryCouponProductRelationListData
	QueryFlashPromotionDetailReq                  = smsclient.QueryFlashPromotionDetailReq
	QueryFlashPromotionDetailResp                 = smsclient.QueryFlashPromotionDetailResp
	QueryFlashPromotionListByDateReq              = smsclient.QueryFlashPromotionListByDateReq
	QueryFlashPromotionListByDateResp             = smsclient.QueryFlashPromotionListByDateResp
	QueryFlashPromotionListReq                    = smsclient.QueryFlashPromotionListReq
	QueryFlashPromotionListResp                   = smsclient.QueryFlashPromotionListResp
	QueryFlashPromotionLogDetailReq               = smsclient.QueryFlashPromotionLogDetailReq
	QueryFlashPromotionLogDetailResp              = smsclient.QueryFlashPromotionLogDetailResp
	QueryFlashPromotionLogListReq                 = smsclient.QueryFlashPromotionLogListReq
	QueryFlashPromotionLogListResp                = smsclient.QueryFlashPromotionLogListResp
	QueryFlashPromotionProductRelationDetailReq   = smsclient.QueryFlashPromotionProductRelationDetailReq
	QueryFlashPromotionProductRelationDetailResp  = smsclient.QueryFlashPromotionProductRelationDetailResp
	QueryFlashPromotionProductRelationListReq     = smsclient.QueryFlashPromotionProductRelationListReq
	QueryFlashPromotionProductRelationListResp    = smsclient.QueryFlashPromotionProductRelationListResp
	QueryFlashPromotionSessionDetailReq           = smsclient.QueryFlashPromotionSessionDetailReq
	QueryFlashPromotionSessionDetailResp          = smsclient.QueryFlashPromotionSessionDetailResp
	QueryFlashPromotionSessionListByTimeReq       = smsclient.QueryFlashPromotionSessionListByTimeReq
	QueryFlashPromotionSessionListByTimeResp      = smsclient.QueryFlashPromotionSessionListByTimeResp
	QueryFlashPromotionSessionListReq             = smsclient.QueryFlashPromotionSessionListReq
	QueryFlashPromotionSessionListResp            = smsclient.QueryFlashPromotionSessionListResp
	QueryHomeAdvertiseDetailReq                   = smsclient.QueryHomeAdvertiseDetailReq
	QueryHomeAdvertiseDetailResp                  = smsclient.QueryHomeAdvertiseDetailResp
	QueryHomeAdvertiseListReq                     = smsclient.QueryHomeAdvertiseListReq
	QueryHomeAdvertiseListResp                    = smsclient.QueryHomeAdvertiseListResp
	QueryHomeBrandDetailReq                       = smsclient.QueryHomeBrandDetailReq
	QueryHomeBrandDetailResp                      = smsclient.QueryHomeBrandDetailResp
	QueryHomeBrandListReq                         = smsclient.QueryHomeBrandListReq
	QueryHomeBrandListResp                        = smsclient.QueryHomeBrandListResp
	QueryHomeNewProductDetailReq                  = smsclient.QueryHomeNewProductDetailReq
	QueryHomeNewProductDetailResp                 = smsclient.QueryHomeNewProductDetailResp
	QueryHomeNewProductListReq                    = smsclient.QueryHomeNewProductListReq
	QueryHomeNewProductListResp                   = smsclient.QueryHomeNewProductListResp
	QueryHomeRecommendProductDetailReq            = smsclient.QueryHomeRecommendProductDetailReq
	QueryHomeRecommendProductDetailResp           = smsclient.QueryHomeRecommendProductDetailResp
	QueryHomeRecommendProductListReq              = smsclient.QueryHomeRecommendProductListReq
	QueryHomeRecommendProductListResp             = smsclient.QueryHomeRecommendProductListResp
	QueryHomeRecommendSubjectDetailReq            = smsclient.QueryHomeRecommendSubjectDetailReq
	QueryHomeRecommendSubjectDetailResp           = smsclient.QueryHomeRecommendSubjectDetailResp
	QueryHomeRecommendSubjectListReq              = smsclient.QueryHomeRecommendSubjectListReq
	QueryHomeRecommendSubjectListResp             = smsclient.QueryHomeRecommendSubjectListResp
	QueryMemberCouponListReq                      = smsclient.QueryMemberCouponListReq
	QueryMemberCouponListResp                     = smsclient.QueryMemberCouponListResp
	UpdateCouponHistoryReq                        = smsclient.UpdateCouponHistoryReq
	UpdateCouponHistoryResp                       = smsclient.UpdateCouponHistoryResp
	UpdateCouponHistoryStatusReq                  = smsclient.UpdateCouponHistoryStatusReq
	UpdateCouponHistoryStatusResp                 = smsclient.UpdateCouponHistoryStatusResp
	UpdateFlashPromotionProductRelationReq        = smsclient.UpdateFlashPromotionProductRelationReq
	UpdateFlashPromotionProductRelationResp       = smsclient.UpdateFlashPromotionProductRelationResp
	UpdateFlashPromotionReq                       = smsclient.UpdateFlashPromotionReq
	UpdateFlashPromotionResp                      = smsclient.UpdateFlashPromotionResp
	UpdateFlashPromotionSessionReq                = smsclient.UpdateFlashPromotionSessionReq
	UpdateFlashPromotionSessionResp               = smsclient.UpdateFlashPromotionSessionResp
	UpdateFlashPromotionSessionStatusReq          = smsclient.UpdateFlashPromotionSessionStatusReq
	UpdateFlashPromotionSessionStatusResp         = smsclient.UpdateFlashPromotionSessionStatusResp
	UpdateFlashPromotionStatusReq                 = smsclient.UpdateFlashPromotionStatusReq
	UpdateFlashPromotionStatusResp                = smsclient.UpdateFlashPromotionStatusResp
	UpdateHomeAdvertiseReq                        = smsclient.UpdateHomeAdvertiseReq
	UpdateHomeAdvertiseResp                       = smsclient.UpdateHomeAdvertiseResp
	UpdateHomeAdvertiseStatusReq                  = smsclient.UpdateHomeAdvertiseStatusReq
	UpdateHomeAdvertiseStatusResp                 = smsclient.UpdateHomeAdvertiseStatusResp
	UpdateHomeBrandSortReq                        = smsclient.UpdateHomeBrandSortReq
	UpdateHomeBrandSortResp                       = smsclient.UpdateHomeBrandSortResp
	UpdateHomeBrandStatusReq                      = smsclient.UpdateHomeBrandStatusReq
	UpdateHomeBrandStatusResp                     = smsclient.UpdateHomeBrandStatusResp
	UpdateHomeNewProductStatusReq                 = smsclient.UpdateHomeNewProductStatusReq
	UpdateHomeNewProductStatusResp                = smsclient.UpdateHomeNewProductStatusResp
	UpdateHomeRecommendProductStatusReq           = smsclient.UpdateHomeRecommendProductStatusReq
	UpdateHomeRecommendProductStatusResp          = smsclient.UpdateHomeRecommendProductStatusResp
	UpdateHomeRecommendSubjectStatusReq           = smsclient.UpdateHomeRecommendSubjectStatusReq
	UpdateHomeRecommendSubjectStatusResp          = smsclient.UpdateHomeRecommendSubjectStatusResp
	UpdateNewProductSortReq                       = smsclient.UpdateNewProductSortReq
	UpdateNewProductSortResp                      = smsclient.UpdateNewProductSortResp
	UpdateRecommendProductSortReq                 = smsclient.UpdateRecommendProductSortReq
	UpdateRecommendProductSortResp                = smsclient.UpdateRecommendProductSortResp
	UpdateRecommendSubjectSortReq                 = smsclient.UpdateRecommendSubjectSortReq
	UpdateRecommendSubjectSortResp                = smsclient.UpdateRecommendSubjectSortResp

	HomeNewProductService interface {
		// 添加新鲜好物表
		AddHomeNewProduct(ctx context.Context, in *AddHomeNewProductReq, opts ...grpc.CallOption) (*AddHomeNewProductResp, error)
		// 删除新鲜好物表
		DeleteHomeNewProduct(ctx context.Context, in *DeleteHomeNewProductReq, opts ...grpc.CallOption) (*DeleteHomeNewProductResp, error)
		// 修改首页新品排序
		UpdateNewProductSort(ctx context.Context, in *UpdateNewProductSortReq, opts ...grpc.CallOption) (*UpdateNewProductSortResp, error)
		// 更新新鲜好物表状态
		UpdateHomeNewProductStatus(ctx context.Context, in *UpdateHomeNewProductStatusReq, opts ...grpc.CallOption) (*UpdateHomeNewProductStatusResp, error)
		// 查询新鲜好物表详情
		QueryHomeNewProductDetail(ctx context.Context, in *QueryHomeNewProductDetailReq, opts ...grpc.CallOption) (*QueryHomeNewProductDetailResp, error)
		// 查询新鲜好物表列表
		QueryHomeNewProductList(ctx context.Context, in *QueryHomeNewProductListReq, opts ...grpc.CallOption) (*QueryHomeNewProductListResp, error)
	}

	defaultHomeNewProductService struct {
		cli zrpc.Client
	}
)

func NewHomeNewProductService(cli zrpc.Client) HomeNewProductService {
	return &defaultHomeNewProductService{
		cli: cli,
	}
}

// 添加新鲜好物表
func (m *defaultHomeNewProductService) AddHomeNewProduct(ctx context.Context, in *AddHomeNewProductReq, opts ...grpc.CallOption) (*AddHomeNewProductResp, error) {
	client := smsclient.NewHomeNewProductServiceClient(m.cli.Conn())
	return client.AddHomeNewProduct(ctx, in, opts...)
}

// 删除新鲜好物表
func (m *defaultHomeNewProductService) DeleteHomeNewProduct(ctx context.Context, in *DeleteHomeNewProductReq, opts ...grpc.CallOption) (*DeleteHomeNewProductResp, error) {
	client := smsclient.NewHomeNewProductServiceClient(m.cli.Conn())
	return client.DeleteHomeNewProduct(ctx, in, opts...)
}

// 修改首页新品排序
func (m *defaultHomeNewProductService) UpdateNewProductSort(ctx context.Context, in *UpdateNewProductSortReq, opts ...grpc.CallOption) (*UpdateNewProductSortResp, error) {
	client := smsclient.NewHomeNewProductServiceClient(m.cli.Conn())
	return client.UpdateNewProductSort(ctx, in, opts...)
}

// 更新新鲜好物表状态
func (m *defaultHomeNewProductService) UpdateHomeNewProductStatus(ctx context.Context, in *UpdateHomeNewProductStatusReq, opts ...grpc.CallOption) (*UpdateHomeNewProductStatusResp, error) {
	client := smsclient.NewHomeNewProductServiceClient(m.cli.Conn())
	return client.UpdateHomeNewProductStatus(ctx, in, opts...)
}

// 查询新鲜好物表详情
func (m *defaultHomeNewProductService) QueryHomeNewProductDetail(ctx context.Context, in *QueryHomeNewProductDetailReq, opts ...grpc.CallOption) (*QueryHomeNewProductDetailResp, error) {
	client := smsclient.NewHomeNewProductServiceClient(m.cli.Conn())
	return client.QueryHomeNewProductDetail(ctx, in, opts...)
}

// 查询新鲜好物表列表
func (m *defaultHomeNewProductService) QueryHomeNewProductList(ctx context.Context, in *QueryHomeNewProductListReq, opts ...grpc.CallOption) (*QueryHomeNewProductListResp, error) {
	client := smsclient.NewHomeNewProductServiceClient(m.cli.Conn())
	return client.QueryHomeNewProductList(ctx, in, opts...)
}
