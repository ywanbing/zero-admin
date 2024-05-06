// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/feihua/zero-admin/rpc/ums/gen/model"
)

func newUmsMemberProductCategoryRelation(db *gorm.DB, opts ...gen.DOOption) umsMemberProductCategoryRelation {
	_umsMemberProductCategoryRelation := umsMemberProductCategoryRelation{}

	_umsMemberProductCategoryRelation.umsMemberProductCategoryRelationDo.UseDB(db, opts...)
	_umsMemberProductCategoryRelation.umsMemberProductCategoryRelationDo.UseModel(&model.UmsMemberProductCategoryRelation{})

	tableName := _umsMemberProductCategoryRelation.umsMemberProductCategoryRelationDo.TableName()
	_umsMemberProductCategoryRelation.ALL = field.NewAsterisk(tableName)
	_umsMemberProductCategoryRelation.ID = field.NewInt64(tableName, "id")
	_umsMemberProductCategoryRelation.MemberID = field.NewInt64(tableName, "member_id")
	_umsMemberProductCategoryRelation.ProductCategoryID = field.NewInt64(tableName, "product_category_id")

	_umsMemberProductCategoryRelation.fillFieldMap()

	return _umsMemberProductCategoryRelation
}

// umsMemberProductCategoryRelation 会员与产品分类关系表（用户喜欢的分类）
type umsMemberProductCategoryRelation struct {
	umsMemberProductCategoryRelationDo umsMemberProductCategoryRelationDo

	ALL               field.Asterisk
	ID                field.Int64
	MemberID          field.Int64 // 会员id
	ProductCategoryID field.Int64 // 商品分类id

	fieldMap map[string]field.Expr
}

func (u umsMemberProductCategoryRelation) Table(newTableName string) *umsMemberProductCategoryRelation {
	u.umsMemberProductCategoryRelationDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsMemberProductCategoryRelation) As(alias string) *umsMemberProductCategoryRelation {
	u.umsMemberProductCategoryRelationDo.DO = *(u.umsMemberProductCategoryRelationDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsMemberProductCategoryRelation) updateTableName(table string) *umsMemberProductCategoryRelation {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.MemberID = field.NewInt64(table, "member_id")
	u.ProductCategoryID = field.NewInt64(table, "product_category_id")

	u.fillFieldMap()

	return u
}

func (u *umsMemberProductCategoryRelation) WithContext(ctx context.Context) IUmsMemberProductCategoryRelationDo {
	return u.umsMemberProductCategoryRelationDo.WithContext(ctx)
}

func (u umsMemberProductCategoryRelation) TableName() string {
	return u.umsMemberProductCategoryRelationDo.TableName()
}

func (u umsMemberProductCategoryRelation) Alias() string {
	return u.umsMemberProductCategoryRelationDo.Alias()
}

func (u umsMemberProductCategoryRelation) Columns(cols ...field.Expr) gen.Columns {
	return u.umsMemberProductCategoryRelationDo.Columns(cols...)
}

func (u *umsMemberProductCategoryRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsMemberProductCategoryRelation) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 3)
	u.fieldMap["id"] = u.ID
	u.fieldMap["member_id"] = u.MemberID
	u.fieldMap["product_category_id"] = u.ProductCategoryID
}

func (u umsMemberProductCategoryRelation) clone(db *gorm.DB) umsMemberProductCategoryRelation {
	u.umsMemberProductCategoryRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsMemberProductCategoryRelation) replaceDB(db *gorm.DB) umsMemberProductCategoryRelation {
	u.umsMemberProductCategoryRelationDo.ReplaceDB(db)
	return u
}

type umsMemberProductCategoryRelationDo struct{ gen.DO }

type IUmsMemberProductCategoryRelationDo interface {
	gen.SubQuery
	Debug() IUmsMemberProductCategoryRelationDo
	WithContext(ctx context.Context) IUmsMemberProductCategoryRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsMemberProductCategoryRelationDo
	WriteDB() IUmsMemberProductCategoryRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsMemberProductCategoryRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsMemberProductCategoryRelationDo
	Not(conds ...gen.Condition) IUmsMemberProductCategoryRelationDo
	Or(conds ...gen.Condition) IUmsMemberProductCategoryRelationDo
	Select(conds ...field.Expr) IUmsMemberProductCategoryRelationDo
	Where(conds ...gen.Condition) IUmsMemberProductCategoryRelationDo
	Order(conds ...field.Expr) IUmsMemberProductCategoryRelationDo
	Distinct(cols ...field.Expr) IUmsMemberProductCategoryRelationDo
	Omit(cols ...field.Expr) IUmsMemberProductCategoryRelationDo
	Join(table schema.Tabler, on ...field.Expr) IUmsMemberProductCategoryRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberProductCategoryRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberProductCategoryRelationDo
	Group(cols ...field.Expr) IUmsMemberProductCategoryRelationDo
	Having(conds ...gen.Condition) IUmsMemberProductCategoryRelationDo
	Limit(limit int) IUmsMemberProductCategoryRelationDo
	Offset(offset int) IUmsMemberProductCategoryRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberProductCategoryRelationDo
	Unscoped() IUmsMemberProductCategoryRelationDo
	Create(values ...*model.UmsMemberProductCategoryRelation) error
	CreateInBatches(values []*model.UmsMemberProductCategoryRelation, batchSize int) error
	Save(values ...*model.UmsMemberProductCategoryRelation) error
	First() (*model.UmsMemberProductCategoryRelation, error)
	Take() (*model.UmsMemberProductCategoryRelation, error)
	Last() (*model.UmsMemberProductCategoryRelation, error)
	Find() ([]*model.UmsMemberProductCategoryRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberProductCategoryRelation, err error)
	FindInBatches(result *[]*model.UmsMemberProductCategoryRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsMemberProductCategoryRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsMemberProductCategoryRelationDo
	Assign(attrs ...field.AssignExpr) IUmsMemberProductCategoryRelationDo
	Joins(fields ...field.RelationField) IUmsMemberProductCategoryRelationDo
	Preload(fields ...field.RelationField) IUmsMemberProductCategoryRelationDo
	FirstOrInit() (*model.UmsMemberProductCategoryRelation, error)
	FirstOrCreate() (*model.UmsMemberProductCategoryRelation, error)
	FindByPage(offset int, limit int) (result []*model.UmsMemberProductCategoryRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsMemberProductCategoryRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsMemberProductCategoryRelationDo) Debug() IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Debug())
}

func (u umsMemberProductCategoryRelationDo) WithContext(ctx context.Context) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsMemberProductCategoryRelationDo) ReadDB() IUmsMemberProductCategoryRelationDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsMemberProductCategoryRelationDo) WriteDB() IUmsMemberProductCategoryRelationDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsMemberProductCategoryRelationDo) Session(config *gorm.Session) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsMemberProductCategoryRelationDo) Clauses(conds ...clause.Expression) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsMemberProductCategoryRelationDo) Returning(value interface{}, columns ...string) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsMemberProductCategoryRelationDo) Not(conds ...gen.Condition) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsMemberProductCategoryRelationDo) Or(conds ...gen.Condition) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsMemberProductCategoryRelationDo) Select(conds ...field.Expr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsMemberProductCategoryRelationDo) Where(conds ...gen.Condition) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsMemberProductCategoryRelationDo) Order(conds ...field.Expr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsMemberProductCategoryRelationDo) Distinct(cols ...field.Expr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsMemberProductCategoryRelationDo) Omit(cols ...field.Expr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsMemberProductCategoryRelationDo) Join(table schema.Tabler, on ...field.Expr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsMemberProductCategoryRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsMemberProductCategoryRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsMemberProductCategoryRelationDo) Group(cols ...field.Expr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsMemberProductCategoryRelationDo) Having(conds ...gen.Condition) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsMemberProductCategoryRelationDo) Limit(limit int) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsMemberProductCategoryRelationDo) Offset(offset int) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsMemberProductCategoryRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsMemberProductCategoryRelationDo) Unscoped() IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsMemberProductCategoryRelationDo) Create(values ...*model.UmsMemberProductCategoryRelation) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsMemberProductCategoryRelationDo) CreateInBatches(values []*model.UmsMemberProductCategoryRelation, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsMemberProductCategoryRelationDo) Save(values ...*model.UmsMemberProductCategoryRelation) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsMemberProductCategoryRelationDo) First() (*model.UmsMemberProductCategoryRelation, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberProductCategoryRelation), nil
	}
}

func (u umsMemberProductCategoryRelationDo) Take() (*model.UmsMemberProductCategoryRelation, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberProductCategoryRelation), nil
	}
}

func (u umsMemberProductCategoryRelationDo) Last() (*model.UmsMemberProductCategoryRelation, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberProductCategoryRelation), nil
	}
}

func (u umsMemberProductCategoryRelationDo) Find() ([]*model.UmsMemberProductCategoryRelation, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsMemberProductCategoryRelation), err
}

func (u umsMemberProductCategoryRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberProductCategoryRelation, err error) {
	buf := make([]*model.UmsMemberProductCategoryRelation, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsMemberProductCategoryRelationDo) FindInBatches(result *[]*model.UmsMemberProductCategoryRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsMemberProductCategoryRelationDo) Attrs(attrs ...field.AssignExpr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsMemberProductCategoryRelationDo) Assign(attrs ...field.AssignExpr) IUmsMemberProductCategoryRelationDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsMemberProductCategoryRelationDo) Joins(fields ...field.RelationField) IUmsMemberProductCategoryRelationDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsMemberProductCategoryRelationDo) Preload(fields ...field.RelationField) IUmsMemberProductCategoryRelationDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsMemberProductCategoryRelationDo) FirstOrInit() (*model.UmsMemberProductCategoryRelation, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberProductCategoryRelation), nil
	}
}

func (u umsMemberProductCategoryRelationDo) FirstOrCreate() (*model.UmsMemberProductCategoryRelation, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberProductCategoryRelation), nil
	}
}

func (u umsMemberProductCategoryRelationDo) FindByPage(offset int, limit int) (result []*model.UmsMemberProductCategoryRelation, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u umsMemberProductCategoryRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsMemberProductCategoryRelationDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsMemberProductCategoryRelationDo) Delete(models ...*model.UmsMemberProductCategoryRelation) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsMemberProductCategoryRelationDo) withDO(do gen.Dao) *umsMemberProductCategoryRelationDo {
	u.DO = *do.(*gen.DO)
	return u
}
