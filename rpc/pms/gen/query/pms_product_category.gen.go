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

func newPmsProductCategory(db *gorm.DB, opts ...gen.DOOption) pmsProductCategory {
	_pmsProductCategory := pmsProductCategory{}

	_pmsProductCategory.pmsProductCategoryDo.UseDB(db, opts...)
	_pmsProductCategory.pmsProductCategoryDo.UseModel(&model.PmsProductCategory{})

	tableName := _pmsProductCategory.pmsProductCategoryDo.TableName()
	_pmsProductCategory.ALL = field.NewAsterisk(tableName)
	_pmsProductCategory.ID = field.NewInt64(tableName, "id")
	_pmsProductCategory.ParentID = field.NewInt64(tableName, "parent_id")
	_pmsProductCategory.Name = field.NewString(tableName, "name")
	_pmsProductCategory.Level = field.NewInt32(tableName, "level")
	_pmsProductCategory.ProductCount = field.NewInt32(tableName, "product_count")
	_pmsProductCategory.ProductUnit = field.NewString(tableName, "product_unit")
	_pmsProductCategory.NavStatus = field.NewInt32(tableName, "nav_status")
	_pmsProductCategory.ShowStatus = field.NewInt32(tableName, "show_status")
	_pmsProductCategory.Sort = field.NewInt32(tableName, "sort")
	_pmsProductCategory.Icon = field.NewString(tableName, "icon")
	_pmsProductCategory.Keywords = field.NewString(tableName, "keywords")
	_pmsProductCategory.Description = field.NewString(tableName, "description")

	_pmsProductCategory.fillFieldMap()

	return _pmsProductCategory
}

// pmsProductCategory 产品分类
type pmsProductCategory struct {
	pmsProductCategoryDo pmsProductCategoryDo

	ALL          field.Asterisk
	ID           field.Int64
	ParentID     field.Int64 // 上机分类的编号：0表示一级分类
	Name         field.String
	Level        field.Int32 // 分类级别：0->1级；1->2级
	ProductCount field.Int32
	ProductUnit  field.String
	NavStatus    field.Int32 // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   field.Int32 // 显示状态：0->不显示；1->显示
	Sort         field.Int32
	Icon         field.String // 图标
	Keywords     field.String
	Description  field.String // 描述

	fieldMap map[string]field.Expr
}

func (p pmsProductCategory) Table(newTableName string) *pmsProductCategory {
	p.pmsProductCategoryDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsProductCategory) As(alias string) *pmsProductCategory {
	p.pmsProductCategoryDo.DO = *(p.pmsProductCategoryDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsProductCategory) updateTableName(table string) *pmsProductCategory {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ParentID = field.NewInt64(table, "parent_id")
	p.Name = field.NewString(table, "name")
	p.Level = field.NewInt32(table, "level")
	p.ProductCount = field.NewInt32(table, "product_count")
	p.ProductUnit = field.NewString(table, "product_unit")
	p.NavStatus = field.NewInt32(table, "nav_status")
	p.ShowStatus = field.NewInt32(table, "show_status")
	p.Sort = field.NewInt32(table, "sort")
	p.Icon = field.NewString(table, "icon")
	p.Keywords = field.NewString(table, "keywords")
	p.Description = field.NewString(table, "description")

	p.fillFieldMap()

	return p
}

func (p *pmsProductCategory) WithContext(ctx context.Context) IPmsProductCategoryDo {
	return p.pmsProductCategoryDo.WithContext(ctx)
}

func (p pmsProductCategory) TableName() string { return p.pmsProductCategoryDo.TableName() }

func (p pmsProductCategory) Alias() string { return p.pmsProductCategoryDo.Alias() }

func (p pmsProductCategory) Columns(cols ...field.Expr) gen.Columns {
	return p.pmsProductCategoryDo.Columns(cols...)
}

func (p *pmsProductCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsProductCategory) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["parent_id"] = p.ParentID
	p.fieldMap["name"] = p.Name
	p.fieldMap["level"] = p.Level
	p.fieldMap["product_count"] = p.ProductCount
	p.fieldMap["product_unit"] = p.ProductUnit
	p.fieldMap["nav_status"] = p.NavStatus
	p.fieldMap["show_status"] = p.ShowStatus
	p.fieldMap["sort"] = p.Sort
	p.fieldMap["icon"] = p.Icon
	p.fieldMap["keywords"] = p.Keywords
	p.fieldMap["description"] = p.Description
}

func (p pmsProductCategory) clone(db *gorm.DB) pmsProductCategory {
	p.pmsProductCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsProductCategory) replaceDB(db *gorm.DB) pmsProductCategory {
	p.pmsProductCategoryDo.ReplaceDB(db)
	return p
}

type pmsProductCategoryDo struct{ gen.DO }

type IPmsProductCategoryDo interface {
	gen.SubQuery
	Debug() IPmsProductCategoryDo
	WithContext(ctx context.Context) IPmsProductCategoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsProductCategoryDo
	WriteDB() IPmsProductCategoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsProductCategoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsProductCategoryDo
	Not(conds ...gen.Condition) IPmsProductCategoryDo
	Or(conds ...gen.Condition) IPmsProductCategoryDo
	Select(conds ...field.Expr) IPmsProductCategoryDo
	Where(conds ...gen.Condition) IPmsProductCategoryDo
	Order(conds ...field.Expr) IPmsProductCategoryDo
	Distinct(cols ...field.Expr) IPmsProductCategoryDo
	Omit(cols ...field.Expr) IPmsProductCategoryDo
	Join(table schema.Tabler, on ...field.Expr) IPmsProductCategoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductCategoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductCategoryDo
	Group(cols ...field.Expr) IPmsProductCategoryDo
	Having(conds ...gen.Condition) IPmsProductCategoryDo
	Limit(limit int) IPmsProductCategoryDo
	Offset(offset int) IPmsProductCategoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductCategoryDo
	Unscoped() IPmsProductCategoryDo
	Create(values ...*model.PmsProductCategory) error
	CreateInBatches(values []*model.PmsProductCategory, batchSize int) error
	Save(values ...*model.PmsProductCategory) error
	First() (*model.PmsProductCategory, error)
	Take() (*model.PmsProductCategory, error)
	Last() (*model.PmsProductCategory, error)
	Find() ([]*model.PmsProductCategory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductCategory, err error)
	FindInBatches(result *[]*model.PmsProductCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsProductCategory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsProductCategoryDo
	Assign(attrs ...field.AssignExpr) IPmsProductCategoryDo
	Joins(fields ...field.RelationField) IPmsProductCategoryDo
	Preload(fields ...field.RelationField) IPmsProductCategoryDo
	FirstOrInit() (*model.PmsProductCategory, error)
	FirstOrCreate() (*model.PmsProductCategory, error)
	FindByPage(offset int, limit int) (result []*model.PmsProductCategory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsProductCategoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsProductCategoryDo) Debug() IPmsProductCategoryDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsProductCategoryDo) WithContext(ctx context.Context) IPmsProductCategoryDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsProductCategoryDo) ReadDB() IPmsProductCategoryDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsProductCategoryDo) WriteDB() IPmsProductCategoryDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsProductCategoryDo) Session(config *gorm.Session) IPmsProductCategoryDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsProductCategoryDo) Clauses(conds ...clause.Expression) IPmsProductCategoryDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsProductCategoryDo) Returning(value interface{}, columns ...string) IPmsProductCategoryDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsProductCategoryDo) Not(conds ...gen.Condition) IPmsProductCategoryDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsProductCategoryDo) Or(conds ...gen.Condition) IPmsProductCategoryDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsProductCategoryDo) Select(conds ...field.Expr) IPmsProductCategoryDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsProductCategoryDo) Where(conds ...gen.Condition) IPmsProductCategoryDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsProductCategoryDo) Order(conds ...field.Expr) IPmsProductCategoryDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsProductCategoryDo) Distinct(cols ...field.Expr) IPmsProductCategoryDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsProductCategoryDo) Omit(cols ...field.Expr) IPmsProductCategoryDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsProductCategoryDo) Join(table schema.Tabler, on ...field.Expr) IPmsProductCategoryDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsProductCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductCategoryDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsProductCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductCategoryDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsProductCategoryDo) Group(cols ...field.Expr) IPmsProductCategoryDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsProductCategoryDo) Having(conds ...gen.Condition) IPmsProductCategoryDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsProductCategoryDo) Limit(limit int) IPmsProductCategoryDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsProductCategoryDo) Offset(offset int) IPmsProductCategoryDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsProductCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductCategoryDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsProductCategoryDo) Unscoped() IPmsProductCategoryDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsProductCategoryDo) Create(values ...*model.PmsProductCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsProductCategoryDo) CreateInBatches(values []*model.PmsProductCategory, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsProductCategoryDo) Save(values ...*model.PmsProductCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsProductCategoryDo) First() (*model.PmsProductCategory, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategory), nil
	}
}

func (p pmsProductCategoryDo) Take() (*model.PmsProductCategory, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategory), nil
	}
}

func (p pmsProductCategoryDo) Last() (*model.PmsProductCategory, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategory), nil
	}
}

func (p pmsProductCategoryDo) Find() ([]*model.PmsProductCategory, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsProductCategory), err
}

func (p pmsProductCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductCategory, err error) {
	buf := make([]*model.PmsProductCategory, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsProductCategoryDo) FindInBatches(result *[]*model.PmsProductCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsProductCategoryDo) Attrs(attrs ...field.AssignExpr) IPmsProductCategoryDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsProductCategoryDo) Assign(attrs ...field.AssignExpr) IPmsProductCategoryDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsProductCategoryDo) Joins(fields ...field.RelationField) IPmsProductCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsProductCategoryDo) Preload(fields ...field.RelationField) IPmsProductCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsProductCategoryDo) FirstOrInit() (*model.PmsProductCategory, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategory), nil
	}
}

func (p pmsProductCategoryDo) FirstOrCreate() (*model.PmsProductCategory, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductCategory), nil
	}
}

func (p pmsProductCategoryDo) FindByPage(offset int, limit int) (result []*model.PmsProductCategory, count int64, err error) {
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

func (p pmsProductCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsProductCategoryDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsProductCategoryDo) Delete(models ...*model.PmsProductCategory) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsProductCategoryDo) withDO(do gen.Dao) *pmsProductCategoryDo {
	p.DO = *do.(*gen.DO)
	return p
}
