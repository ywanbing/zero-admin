// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                                = new(Query)
	CmsSubjectProductRelation        *cmsSubjectProductRelation
	UmsGrowthChangeHistory           *umsGrowthChangeHistory
	UmsIntegrationChangeHistory      *umsIntegrationChangeHistory
	UmsIntegrationConsumeSetting     *umsIntegrationConsumeSetting
	UmsMember                        *umsMember
	UmsMemberBrandAttention          *umsMemberBrandAttention
	UmsMemberLevel                   *umsMemberLevel
	UmsMemberLoginLog                *umsMemberLoginLog
	UmsMemberMemberTagRelation       *umsMemberMemberTagRelation
	UmsMemberProductCategoryRelation *umsMemberProductCategoryRelation
	UmsMemberProductCollection       *umsMemberProductCollection
	UmsMemberReadHistory             *umsMemberReadHistory
	UmsMemberReceiveAddress          *umsMemberReceiveAddress
	UmsMemberRuleSetting             *umsMemberRuleSetting
	UmsMemberStatisticsInfo          *umsMemberStatisticsInfo
	UmsMemberTag                     *umsMemberTag
	UmsMemberTask                    *umsMemberTask
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CmsSubjectProductRelation = &Q.CmsSubjectProductRelation
	UmsGrowthChangeHistory = &Q.UmsGrowthChangeHistory
	UmsIntegrationChangeHistory = &Q.UmsIntegrationChangeHistory
	UmsIntegrationConsumeSetting = &Q.UmsIntegrationConsumeSetting
	UmsMember = &Q.UmsMember
	UmsMemberBrandAttention = &Q.UmsMemberBrandAttention
	UmsMemberLevel = &Q.UmsMemberLevel
	UmsMemberLoginLog = &Q.UmsMemberLoginLog
	UmsMemberMemberTagRelation = &Q.UmsMemberMemberTagRelation
	UmsMemberProductCategoryRelation = &Q.UmsMemberProductCategoryRelation
	UmsMemberProductCollection = &Q.UmsMemberProductCollection
	UmsMemberReadHistory = &Q.UmsMemberReadHistory
	UmsMemberReceiveAddress = &Q.UmsMemberReceiveAddress
	UmsMemberRuleSetting = &Q.UmsMemberRuleSetting
	UmsMemberStatisticsInfo = &Q.UmsMemberStatisticsInfo
	UmsMemberTag = &Q.UmsMemberTag
	UmsMemberTask = &Q.UmsMemberTask
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                               db,
		CmsSubjectProductRelation:        newCmsSubjectProductRelation(db, opts...),
		UmsGrowthChangeHistory:           newUmsGrowthChangeHistory(db, opts...),
		UmsIntegrationChangeHistory:      newUmsIntegrationChangeHistory(db, opts...),
		UmsIntegrationConsumeSetting:     newUmsIntegrationConsumeSetting(db, opts...),
		UmsMember:                        newUmsMember(db, opts...),
		UmsMemberBrandAttention:          newUmsMemberBrandAttention(db, opts...),
		UmsMemberLevel:                   newUmsMemberLevel(db, opts...),
		UmsMemberLoginLog:                newUmsMemberLoginLog(db, opts...),
		UmsMemberMemberTagRelation:       newUmsMemberMemberTagRelation(db, opts...),
		UmsMemberProductCategoryRelation: newUmsMemberProductCategoryRelation(db, opts...),
		UmsMemberProductCollection:       newUmsMemberProductCollection(db, opts...),
		UmsMemberReadHistory:             newUmsMemberReadHistory(db, opts...),
		UmsMemberReceiveAddress:          newUmsMemberReceiveAddress(db, opts...),
		UmsMemberRuleSetting:             newUmsMemberRuleSetting(db, opts...),
		UmsMemberStatisticsInfo:          newUmsMemberStatisticsInfo(db, opts...),
		UmsMemberTag:                     newUmsMemberTag(db, opts...),
		UmsMemberTask:                    newUmsMemberTask(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CmsSubjectProductRelation        cmsSubjectProductRelation
	UmsGrowthChangeHistory           umsGrowthChangeHistory
	UmsIntegrationChangeHistory      umsIntegrationChangeHistory
	UmsIntegrationConsumeSetting     umsIntegrationConsumeSetting
	UmsMember                        umsMember
	UmsMemberBrandAttention          umsMemberBrandAttention
	UmsMemberLevel                   umsMemberLevel
	UmsMemberLoginLog                umsMemberLoginLog
	UmsMemberMemberTagRelation       umsMemberMemberTagRelation
	UmsMemberProductCategoryRelation umsMemberProductCategoryRelation
	UmsMemberProductCollection       umsMemberProductCollection
	UmsMemberReadHistory             umsMemberReadHistory
	UmsMemberReceiveAddress          umsMemberReceiveAddress
	UmsMemberRuleSetting             umsMemberRuleSetting
	UmsMemberStatisticsInfo          umsMemberStatisticsInfo
	UmsMemberTag                     umsMemberTag
	UmsMemberTask                    umsMemberTask
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                               db,
		CmsSubjectProductRelation:        q.CmsSubjectProductRelation.clone(db),
		UmsGrowthChangeHistory:           q.UmsGrowthChangeHistory.clone(db),
		UmsIntegrationChangeHistory:      q.UmsIntegrationChangeHistory.clone(db),
		UmsIntegrationConsumeSetting:     q.UmsIntegrationConsumeSetting.clone(db),
		UmsMember:                        q.UmsMember.clone(db),
		UmsMemberBrandAttention:          q.UmsMemberBrandAttention.clone(db),
		UmsMemberLevel:                   q.UmsMemberLevel.clone(db),
		UmsMemberLoginLog:                q.UmsMemberLoginLog.clone(db),
		UmsMemberMemberTagRelation:       q.UmsMemberMemberTagRelation.clone(db),
		UmsMemberProductCategoryRelation: q.UmsMemberProductCategoryRelation.clone(db),
		UmsMemberProductCollection:       q.UmsMemberProductCollection.clone(db),
		UmsMemberReadHistory:             q.UmsMemberReadHistory.clone(db),
		UmsMemberReceiveAddress:          q.UmsMemberReceiveAddress.clone(db),
		UmsMemberRuleSetting:             q.UmsMemberRuleSetting.clone(db),
		UmsMemberStatisticsInfo:          q.UmsMemberStatisticsInfo.clone(db),
		UmsMemberTag:                     q.UmsMemberTag.clone(db),
		UmsMemberTask:                    q.UmsMemberTask.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                               db,
		CmsSubjectProductRelation:        q.CmsSubjectProductRelation.replaceDB(db),
		UmsGrowthChangeHistory:           q.UmsGrowthChangeHistory.replaceDB(db),
		UmsIntegrationChangeHistory:      q.UmsIntegrationChangeHistory.replaceDB(db),
		UmsIntegrationConsumeSetting:     q.UmsIntegrationConsumeSetting.replaceDB(db),
		UmsMember:                        q.UmsMember.replaceDB(db),
		UmsMemberBrandAttention:          q.UmsMemberBrandAttention.replaceDB(db),
		UmsMemberLevel:                   q.UmsMemberLevel.replaceDB(db),
		UmsMemberLoginLog:                q.UmsMemberLoginLog.replaceDB(db),
		UmsMemberMemberTagRelation:       q.UmsMemberMemberTagRelation.replaceDB(db),
		UmsMemberProductCategoryRelation: q.UmsMemberProductCategoryRelation.replaceDB(db),
		UmsMemberProductCollection:       q.UmsMemberProductCollection.replaceDB(db),
		UmsMemberReadHistory:             q.UmsMemberReadHistory.replaceDB(db),
		UmsMemberReceiveAddress:          q.UmsMemberReceiveAddress.replaceDB(db),
		UmsMemberRuleSetting:             q.UmsMemberRuleSetting.replaceDB(db),
		UmsMemberStatisticsInfo:          q.UmsMemberStatisticsInfo.replaceDB(db),
		UmsMemberTag:                     q.UmsMemberTag.replaceDB(db),
		UmsMemberTask:                    q.UmsMemberTask.replaceDB(db),
	}
}

type queryCtx struct {
	CmsSubjectProductRelation        ICmsSubjectProductRelationDo
	UmsGrowthChangeHistory           IUmsGrowthChangeHistoryDo
	UmsIntegrationChangeHistory      IUmsIntegrationChangeHistoryDo
	UmsIntegrationConsumeSetting     IUmsIntegrationConsumeSettingDo
	UmsMember                        IUmsMemberDo
	UmsMemberBrandAttention          IUmsMemberBrandAttentionDo
	UmsMemberLevel                   IUmsMemberLevelDo
	UmsMemberLoginLog                IUmsMemberLoginLogDo
	UmsMemberMemberTagRelation       IUmsMemberMemberTagRelationDo
	UmsMemberProductCategoryRelation IUmsMemberProductCategoryRelationDo
	UmsMemberProductCollection       IUmsMemberProductCollectionDo
	UmsMemberReadHistory             IUmsMemberReadHistoryDo
	UmsMemberReceiveAddress          IUmsMemberReceiveAddressDo
	UmsMemberRuleSetting             IUmsMemberRuleSettingDo
	UmsMemberStatisticsInfo          IUmsMemberStatisticsInfoDo
	UmsMemberTag                     IUmsMemberTagDo
	UmsMemberTask                    IUmsMemberTaskDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CmsSubjectProductRelation:        q.CmsSubjectProductRelation.WithContext(ctx),
		UmsGrowthChangeHistory:           q.UmsGrowthChangeHistory.WithContext(ctx),
		UmsIntegrationChangeHistory:      q.UmsIntegrationChangeHistory.WithContext(ctx),
		UmsIntegrationConsumeSetting:     q.UmsIntegrationConsumeSetting.WithContext(ctx),
		UmsMember:                        q.UmsMember.WithContext(ctx),
		UmsMemberBrandAttention:          q.UmsMemberBrandAttention.WithContext(ctx),
		UmsMemberLevel:                   q.UmsMemberLevel.WithContext(ctx),
		UmsMemberLoginLog:                q.UmsMemberLoginLog.WithContext(ctx),
		UmsMemberMemberTagRelation:       q.UmsMemberMemberTagRelation.WithContext(ctx),
		UmsMemberProductCategoryRelation: q.UmsMemberProductCategoryRelation.WithContext(ctx),
		UmsMemberProductCollection:       q.UmsMemberProductCollection.WithContext(ctx),
		UmsMemberReadHistory:             q.UmsMemberReadHistory.WithContext(ctx),
		UmsMemberReceiveAddress:          q.UmsMemberReceiveAddress.WithContext(ctx),
		UmsMemberRuleSetting:             q.UmsMemberRuleSetting.WithContext(ctx),
		UmsMemberStatisticsInfo:          q.UmsMemberStatisticsInfo.WithContext(ctx),
		UmsMemberTag:                     q.UmsMemberTag.WithContext(ctx),
		UmsMemberTask:                    q.UmsMemberTask.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
