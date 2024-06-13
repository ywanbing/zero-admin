// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package flashpromotionsessionservice

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

	FlashPromotionSessionService interface {
		// 添加限时购场次表
		AddFlashPromotionSession(ctx context.Context, in *AddFlashPromotionSessionReq, opts ...grpc.CallOption) (*AddFlashPromotionSessionResp, error)
		// 删除限时购场次表
		DeleteFlashPromotionSession(ctx context.Context, in *DeleteFlashPromotionSessionReq, opts ...grpc.CallOption) (*DeleteFlashPromotionSessionResp, error)
		// 更新限时购场次表
		UpdateFlashPromotionSession(ctx context.Context, in *UpdateFlashPromotionSessionReq, opts ...grpc.CallOption) (*UpdateFlashPromotionSessionResp, error)
		// 更新限时购场次表状态
		UpdateFlashPromotionSessionStatus(ctx context.Context, in *UpdateFlashPromotionSessionStatusReq, opts ...grpc.CallOption) (*UpdateFlashPromotionSessionStatusResp, error)
		// 查询限时购场次表详情
		QueryFlashPromotionSessionDetail(ctx context.Context, in *QueryFlashPromotionSessionDetailReq, opts ...grpc.CallOption) (*QueryFlashPromotionSessionDetailResp, error)
		// 查询限时购场次表列表
		QueryFlashPromotionSessionList(ctx context.Context, in *QueryFlashPromotionSessionListReq, opts ...grpc.CallOption) (*QueryFlashPromotionSessionListResp, error)
		// 根据时间查询限时购场次
		QueryFlashPromotionSessionListByTime(ctx context.Context, in *QueryFlashPromotionSessionListByTimeReq, opts ...grpc.CallOption) (*QueryFlashPromotionSessionListByTimeResp, error)
	}

	defaultFlashPromotionSessionService struct {
		cli zrpc.Client
	}
)

func NewFlashPromotionSessionService(cli zrpc.Client) FlashPromotionSessionService {
	return &defaultFlashPromotionSessionService{
		cli: cli,
	}
}

// 添加限时购场次表
func (m *defaultFlashPromotionSessionService) AddFlashPromotionSession(ctx context.Context, in *AddFlashPromotionSessionReq, opts ...grpc.CallOption) (*AddFlashPromotionSessionResp, error) {
	client := smsclient.NewFlashPromotionSessionServiceClient(m.cli.Conn())
	return client.AddFlashPromotionSession(ctx, in, opts...)
}

// 删除限时购场次表
func (m *defaultFlashPromotionSessionService) DeleteFlashPromotionSession(ctx context.Context, in *DeleteFlashPromotionSessionReq, opts ...grpc.CallOption) (*DeleteFlashPromotionSessionResp, error) {
	client := smsclient.NewFlashPromotionSessionServiceClient(m.cli.Conn())
	return client.DeleteFlashPromotionSession(ctx, in, opts...)
}

// 更新限时购场次表
func (m *defaultFlashPromotionSessionService) UpdateFlashPromotionSession(ctx context.Context, in *UpdateFlashPromotionSessionReq, opts ...grpc.CallOption) (*UpdateFlashPromotionSessionResp, error) {
	client := smsclient.NewFlashPromotionSessionServiceClient(m.cli.Conn())
	return client.UpdateFlashPromotionSession(ctx, in, opts...)
}

// 更新限时购场次表状态
func (m *defaultFlashPromotionSessionService) UpdateFlashPromotionSessionStatus(ctx context.Context, in *UpdateFlashPromotionSessionStatusReq, opts ...grpc.CallOption) (*UpdateFlashPromotionSessionStatusResp, error) {
	client := smsclient.NewFlashPromotionSessionServiceClient(m.cli.Conn())
	return client.UpdateFlashPromotionSessionStatus(ctx, in, opts...)
}

// 查询限时购场次表详情
func (m *defaultFlashPromotionSessionService) QueryFlashPromotionSessionDetail(ctx context.Context, in *QueryFlashPromotionSessionDetailReq, opts ...grpc.CallOption) (*QueryFlashPromotionSessionDetailResp, error) {
	client := smsclient.NewFlashPromotionSessionServiceClient(m.cli.Conn())
	return client.QueryFlashPromotionSessionDetail(ctx, in, opts...)
}

// 查询限时购场次表列表
func (m *defaultFlashPromotionSessionService) QueryFlashPromotionSessionList(ctx context.Context, in *QueryFlashPromotionSessionListReq, opts ...grpc.CallOption) (*QueryFlashPromotionSessionListResp, error) {
	client := smsclient.NewFlashPromotionSessionServiceClient(m.cli.Conn())
	return client.QueryFlashPromotionSessionList(ctx, in, opts...)
}

// 根据时间查询限时购场次
func (m *defaultFlashPromotionSessionService) QueryFlashPromotionSessionListByTime(ctx context.Context, in *QueryFlashPromotionSessionListByTimeReq, opts ...grpc.CallOption) (*QueryFlashPromotionSessionListByTimeResp, error) {
	client := smsclient.NewFlashPromotionSessionServiceClient(m.cli.Conn())
	return client.QueryFlashPromotionSessionListByTime(ctx, in, opts...)
}
