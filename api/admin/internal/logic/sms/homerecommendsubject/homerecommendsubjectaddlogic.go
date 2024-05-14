package homerecommendsubject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendSubjectAddLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:42
*/
type HomeRecommendSubjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectAddLogic {
	return HomeRecommendSubjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeRecommendSubjectAdd 添加人气推荐专题
func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(req types.AddHomeRecommendSubjectReq) (*types.AddHomeRecommendSubjectResp, error) {
	listResp, _ := l.svcCtx.SubjectService.SubjectListByIds(l.ctx, &cmsclient.SubjectListByIdsReq{
		Ids: req.SubjectIds,
	})

	var list []*smsclient.HomeRecommendSubjectAddData

	for _, item := range listResp.List {
		list = append(list, &smsclient.HomeRecommendSubjectAddData{
			SubjectId:       item.Id,
			SubjectName:     item.Title,
			RecommendStatus: item.RecommendStatus,
			Sort:            1,
		})
	}

	_, err := l.svcCtx.HomeRecommendSubjectService.HomeRecommendSubjectAdd(l.ctx, &smsclient.HomeRecommendSubjectAddReq{
		RecommendSubjectAddData: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加人气推荐专题信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加人气推荐专题失败")
	}

	return &types.AddHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "添加人气推荐专题成功",
	}, nil
}
