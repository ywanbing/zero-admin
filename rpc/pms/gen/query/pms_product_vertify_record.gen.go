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

func newPmsProductVertifyRecord(db *gorm.DB, opts ...gen.DOOption) pmsProductVertifyRecord {
	_pmsProductVertifyRecord := pmsProductVertifyRecord{}

	_pmsProductVertifyRecord.pmsProductVertifyRecordDo.UseDB(db, opts...)
	_pmsProductVertifyRecord.pmsProductVertifyRecordDo.UseModel(&model.PmsProductVertifyRecord{})

	tableName := _pmsProductVertifyRecord.pmsProductVertifyRecordDo.TableName()
	_pmsProductVertifyRecord.ALL = field.NewAsterisk(tableName)
	_pmsProductVertifyRecord.ID = field.NewInt64(tableName, "id")
	_pmsProductVertifyRecord.ProductID = field.NewInt64(tableName, "product_id")
	_pmsProductVertifyRecord.CreateTime = field.NewTime(tableName, "create_time")
	_pmsProductVertifyRecord.VertifyMan = field.NewString(tableName, "vertify_man")
	_pmsProductVertifyRecord.Status = field.NewInt32(tableName, "status")
	_pmsProductVertifyRecord.Detail = field.NewString(tableName, "detail")

	_pmsProductVertifyRecord.fillFieldMap()

	return _pmsProductVertifyRecord
}

// pmsProductVertifyRecord 商品审核记录
type pmsProductVertifyRecord struct {
	pmsProductVertifyRecordDo pmsProductVertifyRecordDo

	ALL        field.Asterisk
	ID         field.Int64
	ProductID  field.Int64
	CreateTime field.Time
	VertifyMan field.String // 审核人
	Status     field.Int32
	Detail     field.String // 反馈详情

	fieldMap map[string]field.Expr
}

func (p pmsProductVertifyRecord) Table(newTableName string) *pmsProductVertifyRecord {
	p.pmsProductVertifyRecordDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pmsProductVertifyRecord) As(alias string) *pmsProductVertifyRecord {
	p.pmsProductVertifyRecordDo.DO = *(p.pmsProductVertifyRecordDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pmsProductVertifyRecord) updateTableName(table string) *pmsProductVertifyRecord {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ProductID = field.NewInt64(table, "product_id")
	p.CreateTime = field.NewTime(table, "create_time")
	p.VertifyMan = field.NewString(table, "vertify_man")
	p.Status = field.NewInt32(table, "status")
	p.Detail = field.NewString(table, "detail")

	p.fillFieldMap()

	return p
}

func (p *pmsProductVertifyRecord) WithContext(ctx context.Context) IPmsProductVertifyRecordDo {
	return p.pmsProductVertifyRecordDo.WithContext(ctx)
}

func (p pmsProductVertifyRecord) TableName() string { return p.pmsProductVertifyRecordDo.TableName() }

func (p pmsProductVertifyRecord) Alias() string { return p.pmsProductVertifyRecordDo.Alias() }

func (p pmsProductVertifyRecord) Columns(cols ...field.Expr) gen.Columns {
	return p.pmsProductVertifyRecordDo.Columns(cols...)
}

func (p *pmsProductVertifyRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pmsProductVertifyRecord) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["product_id"] = p.ProductID
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["vertify_man"] = p.VertifyMan
	p.fieldMap["status"] = p.Status
	p.fieldMap["detail"] = p.Detail
}

func (p pmsProductVertifyRecord) clone(db *gorm.DB) pmsProductVertifyRecord {
	p.pmsProductVertifyRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pmsProductVertifyRecord) replaceDB(db *gorm.DB) pmsProductVertifyRecord {
	p.pmsProductVertifyRecordDo.ReplaceDB(db)
	return p
}

type pmsProductVertifyRecordDo struct{ gen.DO }

type IPmsProductVertifyRecordDo interface {
	gen.SubQuery
	Debug() IPmsProductVertifyRecordDo
	WithContext(ctx context.Context) IPmsProductVertifyRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPmsProductVertifyRecordDo
	WriteDB() IPmsProductVertifyRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPmsProductVertifyRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPmsProductVertifyRecordDo
	Not(conds ...gen.Condition) IPmsProductVertifyRecordDo
	Or(conds ...gen.Condition) IPmsProductVertifyRecordDo
	Select(conds ...field.Expr) IPmsProductVertifyRecordDo
	Where(conds ...gen.Condition) IPmsProductVertifyRecordDo
	Order(conds ...field.Expr) IPmsProductVertifyRecordDo
	Distinct(cols ...field.Expr) IPmsProductVertifyRecordDo
	Omit(cols ...field.Expr) IPmsProductVertifyRecordDo
	Join(table schema.Tabler, on ...field.Expr) IPmsProductVertifyRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductVertifyRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductVertifyRecordDo
	Group(cols ...field.Expr) IPmsProductVertifyRecordDo
	Having(conds ...gen.Condition) IPmsProductVertifyRecordDo
	Limit(limit int) IPmsProductVertifyRecordDo
	Offset(offset int) IPmsProductVertifyRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductVertifyRecordDo
	Unscoped() IPmsProductVertifyRecordDo
	Create(values ...*model.PmsProductVertifyRecord) error
	CreateInBatches(values []*model.PmsProductVertifyRecord, batchSize int) error
	Save(values ...*model.PmsProductVertifyRecord) error
	First() (*model.PmsProductVertifyRecord, error)
	Take() (*model.PmsProductVertifyRecord, error)
	Last() (*model.PmsProductVertifyRecord, error)
	Find() ([]*model.PmsProductVertifyRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductVertifyRecord, err error)
	FindInBatches(result *[]*model.PmsProductVertifyRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PmsProductVertifyRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPmsProductVertifyRecordDo
	Assign(attrs ...field.AssignExpr) IPmsProductVertifyRecordDo
	Joins(fields ...field.RelationField) IPmsProductVertifyRecordDo
	Preload(fields ...field.RelationField) IPmsProductVertifyRecordDo
	FirstOrInit() (*model.PmsProductVertifyRecord, error)
	FirstOrCreate() (*model.PmsProductVertifyRecord, error)
	FindByPage(offset int, limit int) (result []*model.PmsProductVertifyRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPmsProductVertifyRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pmsProductVertifyRecordDo) Debug() IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Debug())
}

func (p pmsProductVertifyRecordDo) WithContext(ctx context.Context) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pmsProductVertifyRecordDo) ReadDB() IPmsProductVertifyRecordDo {
	return p.Clauses(dbresolver.Read)
}

func (p pmsProductVertifyRecordDo) WriteDB() IPmsProductVertifyRecordDo {
	return p.Clauses(dbresolver.Write)
}

func (p pmsProductVertifyRecordDo) Session(config *gorm.Session) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Session(config))
}

func (p pmsProductVertifyRecordDo) Clauses(conds ...clause.Expression) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pmsProductVertifyRecordDo) Returning(value interface{}, columns ...string) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pmsProductVertifyRecordDo) Not(conds ...gen.Condition) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pmsProductVertifyRecordDo) Or(conds ...gen.Condition) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pmsProductVertifyRecordDo) Select(conds ...field.Expr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pmsProductVertifyRecordDo) Where(conds ...gen.Condition) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pmsProductVertifyRecordDo) Order(conds ...field.Expr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pmsProductVertifyRecordDo) Distinct(cols ...field.Expr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pmsProductVertifyRecordDo) Omit(cols ...field.Expr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pmsProductVertifyRecordDo) Join(table schema.Tabler, on ...field.Expr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pmsProductVertifyRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pmsProductVertifyRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pmsProductVertifyRecordDo) Group(cols ...field.Expr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pmsProductVertifyRecordDo) Having(conds ...gen.Condition) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pmsProductVertifyRecordDo) Limit(limit int) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pmsProductVertifyRecordDo) Offset(offset int) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pmsProductVertifyRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pmsProductVertifyRecordDo) Unscoped() IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pmsProductVertifyRecordDo) Create(values ...*model.PmsProductVertifyRecord) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pmsProductVertifyRecordDo) CreateInBatches(values []*model.PmsProductVertifyRecord, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pmsProductVertifyRecordDo) Save(values ...*model.PmsProductVertifyRecord) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pmsProductVertifyRecordDo) First() (*model.PmsProductVertifyRecord, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductVertifyRecord), nil
	}
}

func (p pmsProductVertifyRecordDo) Take() (*model.PmsProductVertifyRecord, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductVertifyRecord), nil
	}
}

func (p pmsProductVertifyRecordDo) Last() (*model.PmsProductVertifyRecord, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductVertifyRecord), nil
	}
}

func (p pmsProductVertifyRecordDo) Find() ([]*model.PmsProductVertifyRecord, error) {
	result, err := p.DO.Find()
	return result.([]*model.PmsProductVertifyRecord), err
}

func (p pmsProductVertifyRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PmsProductVertifyRecord, err error) {
	buf := make([]*model.PmsProductVertifyRecord, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pmsProductVertifyRecordDo) FindInBatches(result *[]*model.PmsProductVertifyRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pmsProductVertifyRecordDo) Attrs(attrs ...field.AssignExpr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pmsProductVertifyRecordDo) Assign(attrs ...field.AssignExpr) IPmsProductVertifyRecordDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pmsProductVertifyRecordDo) Joins(fields ...field.RelationField) IPmsProductVertifyRecordDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pmsProductVertifyRecordDo) Preload(fields ...field.RelationField) IPmsProductVertifyRecordDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pmsProductVertifyRecordDo) FirstOrInit() (*model.PmsProductVertifyRecord, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductVertifyRecord), nil
	}
}

func (p pmsProductVertifyRecordDo) FirstOrCreate() (*model.PmsProductVertifyRecord, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PmsProductVertifyRecord), nil
	}
}

func (p pmsProductVertifyRecordDo) FindByPage(offset int, limit int) (result []*model.PmsProductVertifyRecord, count int64, err error) {
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

func (p pmsProductVertifyRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pmsProductVertifyRecordDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pmsProductVertifyRecordDo) Delete(models ...*model.PmsProductVertifyRecord) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pmsProductVertifyRecordDo) withDO(do gen.Dao) *pmsProductVertifyRecordDo {
	p.DO = *do.(*gen.DO)
	return p
}
