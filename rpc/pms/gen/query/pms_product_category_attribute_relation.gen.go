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

	"github.com/feihua/zero-admin/rpc/pms/gen/model"
)

func newPmsProductCategoryAttributeRelation(db *gorm.DB, opts ...gen.DOOption) pmsProductCategoryAttributeRelation {
	_pmsProductCategoryAttributeRelation := pmsProductCategoryAttributeRelation{}

	_pmsProductCategoryAttributeRelation.pmsProductCategoryAttributeRelationDo.UseDB(db, opts...)
	_pmsProductCategoryAttributeRelation.pmsProductCategoryAttributeRelationDo.UseModel(&model.PmsProductCategoryAttributeRelation{})

	tableName := _pmsProductCategoryAttributeRelation.pmsProductCategoryAttributeRelationDo.TableName()
	_pmsProductCategoryAttributeRelation.ALL = field.NewAsterisk(tableName)
	_pmsProductCategoryAttributeRelation.ID = field.NewInt64(tableName, "id")
	_pmsProductCategoryAttributeRelation.ProductCategoryID = field.NewInt64(tableName, "product_category_id")
	_pmsProductCategoryAttributeRelation.ProductAttributeID = field.NewInt64(tableName, "product_attribute_id")

	_pmsProductCategoryAttributeRelation.fillFieldMap()

	return _pmsProductCategoryAttributeRelation
}

// pmsProductCategoryAttributeRelation 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type pmsProductCategoryAttributeRelation struct {
	pmsProductCategoryAttributeRelationDo pmsProductCategoryAttributeRelationDo

	ALL                field.Asterisk
	ID                 field.Int64
	ProductCategoryID  field.Int64
	ProductAttributeID field.Int64

	fieldMap map[string]field.Expr
}

func (p pmsProductCategoryAttributeRelation) Table(newTableName string) *pmsProductCategoryAttributeRelation {
	p.pmsProductCategoryAttributeRelationDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsProductCategoryAttributeRelation) As(alias string) *pmsProductCategoryAttributeRelation {
	p.pmsProductCategoryAttributeRelationDo.DO = *(p.pmsProductCategoryAttributeRelationDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsProductCategoryAttributeRelation) updateTableName(table string) *pmsProductCategoryAttributeRelation {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ProductCategoryID = field.NewInt64(table, "product_category_id")
	p.ProductAttributeID = field.NewInt64(table, "product_attribute_id")

	p.fillFieldMap()

	return p
}

func (p *pmsProductCategoryAttributeRelation) WithContext(ctx context.Context) IPmsProductCategoryAttributeRelationDo {
	return p.pmsProductCategoryAttributeRelationDo.WithContext(ctx)
}

func (p pmsProductCategoryAttributeRelation) TableName() string {
	return p.pmsProductCategoryAttributeRelationDo.TableName()
}

func (p pmsProductCategoryAttributeRelation) Alias() string {
	return p.pmsProductCategoryAttributeRelationDo.Alias()
}

func (p pmsProductCategoryAttributeRelation) Columns(cols ...field.Expr) gen.Columns {
	return p.pmsProductCategoryAttributeRelationDo.Columns(cols...)
}

func (p *pmsProductCategoryAttributeRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsProductCategoryAttributeRelation) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 3)
	p.fieldMap["id"] = p.ID
	p.fieldMap["product_category_id"] = p.ProductCategoryID
	p.fieldMap["product_attribute_id"] = p.ProductAttributeID
}

func (p pmsProductCategoryAttributeRelation) clone(db *gorm.DB) pmsProductCategoryAttributeRelation {
	p.pmsProductCategoryAttributeRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsProductCategoryAttributeRelation) replaceDB(db *gorm.DB) pmsProductCategoryAttributeRelation {
	p.pmsProductCategoryAttributeRelationDo.ReplaceDB(db)
	return p
}

type pmsProductCategoryAttributeRelationDo struct{ gen.DO }

type IPmsProductCategoryAttributeRelationDo interface {
	gen.SubQuery
	Debug() IPmsProductCategoryAttributeRelationDo
	WithContext(ctx context.Context) IPmsProductCategoryAttributeRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsProductCategoryAttributeRelationDo
	WriteDB() IPmsProductCategoryAttributeRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsProductCategoryAttributeRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsProductCategoryAttributeRelationDo
	Not(conds ...gen.Condition) IPmsProductCategoryAttributeRelationDo
	Or(conds ...gen.Condition) IPmsProductCategoryAttributeRelationDo
	Select(conds ...field.Expr) IPmsProductCategoryAttributeRelationDo
	Where(conds ...gen.Condition) IPmsProductCategoryAttributeRelationDo
	Order(conds ...field.Expr) IPmsProductCategoryAttributeRelationDo
	Distinct(cols ...field.Expr) IPmsProductCategoryAttributeRelationDo
	Omit(cols ...field.Expr) IPmsProductCategoryAttributeRelationDo
	Join(table schema.Tabler, on ...field.Expr) IPmsProductCategoryAttributeRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductCategoryAttributeRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductCategoryAttributeRelationDo
	Group(cols ...field.Expr) IPmsProductCategoryAttributeRelationDo
	Having(conds ...gen.Condition) IPmsProductCategoryAttributeRelationDo
	Limit(limit int) IPmsProductCategoryAttributeRelationDo
	Offset(offset int) IPmsProductCategoryAttributeRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductCategoryAttributeRelationDo
	Unscoped() IPmsProductCategoryAttributeRelationDo
	Create(values ...*model.PmsProductCategoryAttributeRelation) error
	CreateInBatches(values []*model.PmsProductCategoryAttributeRelation, batchSize int) error
	Save(values ...*model.PmsProductCategoryAttributeRelation) error
	First() (*model.PmsProductCategoryAttributeRelation, error)
	Take() (*model.PmsProductCategoryAttributeRelation, error)
	Last() (*model.PmsProductCategoryAttributeRelation, error)
	Find() ([]*model.PmsProductCategoryAttributeRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductCategoryAttributeRelation, err error)
	FindInBatches(result *[]*model.PmsProductCategoryAttributeRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsProductCategoryAttributeRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsProductCategoryAttributeRelationDo
	Assign(attrs ...field.AssignExpr) IPmsProductCategoryAttributeRelationDo
	Joins(fields ...field.RelationField) IPmsProductCategoryAttributeRelationDo
	Preload(fields ...field.RelationField) IPmsProductCategoryAttributeRelationDo
	FirstOrInit() (*model.PmsProductCategoryAttributeRelation, error)
	FirstOrCreate() (*model.PmsProductCategoryAttributeRelation, error)
	FindByPage(offset int, limit int) (result []*model.PmsProductCategoryAttributeRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsProductCategoryAttributeRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsProductCategoryAttributeRelationDo) Debug() IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsProductCategoryAttributeRelationDo) WithContext(ctx context.Context) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsProductCategoryAttributeRelationDo) ReadDB() IPmsProductCategoryAttributeRelationDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsProductCategoryAttributeRelationDo) WriteDB() IPmsProductCategoryAttributeRelationDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsProductCategoryAttributeRelationDo) Session(config *gorm.Session) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsProductCategoryAttributeRelationDo) Clauses(conds ...clause.Expression) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsProductCategoryAttributeRelationDo) Returning(value interface{}, columns ...string) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsProductCategoryAttributeRelationDo) Not(conds ...gen.Condition) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsProductCategoryAttributeRelationDo) Or(conds ...gen.Condition) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsProductCategoryAttributeRelationDo) Select(conds ...field.Expr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsProductCategoryAttributeRelationDo) Where(conds ...gen.Condition) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsProductCategoryAttributeRelationDo) Order(conds ...field.Expr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsProductCategoryAttributeRelationDo) Distinct(cols ...field.Expr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsProductCategoryAttributeRelationDo) Omit(cols ...field.Expr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsProductCategoryAttributeRelationDo) Join(table schema.Tabler, on ...field.Expr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsProductCategoryAttributeRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsProductCategoryAttributeRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsProductCategoryAttributeRelationDo) Group(cols ...field.Expr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsProductCategoryAttributeRelationDo) Having(conds ...gen.Condition) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsProductCategoryAttributeRelationDo) Limit(limit int) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsProductCategoryAttributeRelationDo) Offset(offset int) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsProductCategoryAttributeRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsProductCategoryAttributeRelationDo) Unscoped() IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsProductCategoryAttributeRelationDo) Create(values ...*model.PmsProductCategoryAttributeRelation) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsProductCategoryAttributeRelationDo) CreateInBatches(values []*model.PmsProductCategoryAttributeRelation, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsProductCategoryAttributeRelationDo) Save(values ...*model.PmsProductCategoryAttributeRelation) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsProductCategoryAttributeRelationDo) First() (*model.PmsProductCategoryAttributeRelation, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategoryAttributeRelation), nil
	}
}

func (p pmsProductCategoryAttributeRelationDo) Take() (*model.PmsProductCategoryAttributeRelation, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategoryAttributeRelation), nil
	}
}

func (p pmsProductCategoryAttributeRelationDo) Last() (*model.PmsProductCategoryAttributeRelation, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategoryAttributeRelation), nil
	}
}

func (p pmsProductCategoryAttributeRelationDo) Find() ([]*model.PmsProductCategoryAttributeRelation, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsProductCategoryAttributeRelation), err
}

func (p pmsProductCategoryAttributeRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductCategoryAttributeRelation, err error) {
	buf := make([]*model.PmsProductCategoryAttributeRelation, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsProductCategoryAttributeRelationDo) FindInBatches(result *[]*model.PmsProductCategoryAttributeRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsProductCategoryAttributeRelationDo) Attrs(attrs ...field.AssignExpr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsProductCategoryAttributeRelationDo) Assign(attrs ...field.AssignExpr) IPmsProductCategoryAttributeRelationDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsProductCategoryAttributeRelationDo) Joins(fields ...field.RelationField) IPmsProductCategoryAttributeRelationDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsProductCategoryAttributeRelationDo) Preload(fields ...field.RelationField) IPmsProductCategoryAttributeRelationDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsProductCategoryAttributeRelationDo) FirstOrInit() (*model.PmsProductCategoryAttributeRelation, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategoryAttributeRelation), nil
	}
}

func (p pmsProductCategoryAttributeRelationDo) FirstOrCreate() (*model.PmsProductCategoryAttributeRelation, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategoryAttributeRelation), nil
	}
}

func (p pmsProductCategoryAttributeRelationDo) FindByPage(offset int, limit int) (result []*model.PmsProductCategoryAttributeRelation, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pmsProductCategoryAttributeRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsProductCategoryAttributeRelationDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsProductCategoryAttributeRelationDo) Delete(models ...*model.PmsProductCategoryAttributeRelation) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsProductCategoryAttributeRelationDo) withDO(do gen.Dao) *pmsProductCategoryAttributeRelationDo {
	p.DO = *do.(*gen.DO)
	return p
}
