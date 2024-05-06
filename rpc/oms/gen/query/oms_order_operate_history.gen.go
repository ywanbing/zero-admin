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

	"github.com/feihua/zero-admin/rpc/oms/gen/model"
)

func newOmsOrderOperateHistory(db *gorm.DB, opts ...gen.DOOption) omsOrderOperateHistory {
	_omsOrderOperateHistory := omsOrderOperateHistory{}

	_omsOrderOperateHistory.omsOrderOperateHistoryDo.UseDB(db, opts...)
	_omsOrderOperateHistory.omsOrderOperateHistoryDo.UseModel(&model.OmsOrderOperateHistory{})

	tableName := _omsOrderOperateHistory.omsOrderOperateHistoryDo.TableName()
	_omsOrderOperateHistory.ALL = field.NewAsterisk(tableName)
	_omsOrderOperateHistory.ID = field.NewInt64(tableName, "id")
	_omsOrderOperateHistory.OrderID = field.NewInt64(tableName, "order_id")
	_omsOrderOperateHistory.OperateMan = field.NewString(tableName, "operate_man")
	_omsOrderOperateHistory.CreateTime = field.NewTime(tableName, "create_time")
	_omsOrderOperateHistory.OrderStatus = field.NewInt32(tableName, "order_status")
	_omsOrderOperateHistory.Note = field.NewString(tableName, "note")

	_omsOrderOperateHistory.fillFieldMap()

	return _omsOrderOperateHistory
}

// omsOrderOperateHistory 订单操作历史记录
type omsOrderOperateHistory struct {
	omsOrderOperateHistoryDo omsOrderOperateHistoryDo

	ALL         field.Asterisk
	ID          field.Int64
	OrderID     field.Int64  // 订单id
	OperateMan  field.String // 操作人：用户；系统；后台管理员
	CreateTime  field.Time   // 操作时间
	OrderStatus field.Int32  // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        field.String // 备注

	fieldMap map[string]field.Expr
}

func (o omsOrderOperateHistory) Table(newTableName string) *omsOrderOperateHistory {
	o.omsOrderOperateHistoryDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o omsOrderOperateHistory) As(alias string) *omsOrderOperateHistory {
	o.omsOrderOperateHistoryDo.DO = *(o.omsOrderOperateHistoryDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *omsOrderOperateHistory) updateTableName(table string) *omsOrderOperateHistory {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.OrderID = field.NewInt64(table, "order_id")
	o.OperateMan = field.NewString(table, "operate_man")
	o.CreateTime = field.NewTime(table, "create_time")
	o.OrderStatus = field.NewInt32(table, "order_status")
	o.Note = field.NewString(table, "note")

	o.fillFieldMap()

	return o
}

func (o *omsOrderOperateHistory) WithContext(ctx context.Context) IOmsOrderOperateHistoryDo {
	return o.omsOrderOperateHistoryDo.WithContext(ctx)
}

func (o omsOrderOperateHistory) TableName() string { return o.omsOrderOperateHistoryDo.TableName() }

func (o omsOrderOperateHistory) Alias() string { return o.omsOrderOperateHistoryDo.Alias() }

func (o omsOrderOperateHistory) Columns(cols ...field.Expr) gen.Columns {
	return o.omsOrderOperateHistoryDo.Columns(cols...)
}

func (o *omsOrderOperateHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *omsOrderOperateHistory) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["id"] = o.ID
	o.fieldMap["order_id"] = o.OrderID
	o.fieldMap["operate_man"] = o.OperateMan
	o.fieldMap["create_time"] = o.CreateTime
	o.fieldMap["order_status"] = o.OrderStatus
	o.fieldMap["note"] = o.Note
}

func (o omsOrderOperateHistory) clone(db *gorm.DB) omsOrderOperateHistory {
	o.omsOrderOperateHistoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o omsOrderOperateHistory) replaceDB(db *gorm.DB) omsOrderOperateHistory {
	o.omsOrderOperateHistoryDo.ReplaceDB(db)
	return o
}

type omsOrderOperateHistoryDo struct{ gen.DO }

type IOmsOrderOperateHistoryDo interface {
	gen.SubQuery
	Debug() IOmsOrderOperateHistoryDo
	WithContext(ctx context.Context) IOmsOrderOperateHistoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOmsOrderOperateHistoryDo
	WriteDB() IOmsOrderOperateHistoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOmsOrderOperateHistoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOmsOrderOperateHistoryDo
	Not(conds ...gen.Condition) IOmsOrderOperateHistoryDo
	Or(conds ...gen.Condition) IOmsOrderOperateHistoryDo
	Select(conds ...field.Expr) IOmsOrderOperateHistoryDo
	Where(conds ...gen.Condition) IOmsOrderOperateHistoryDo
	Order(conds ...field.Expr) IOmsOrderOperateHistoryDo
	Distinct(cols ...field.Expr) IOmsOrderOperateHistoryDo
	Omit(cols ...field.Expr) IOmsOrderOperateHistoryDo
	Join(table schema.Tabler, on ...field.Expr) IOmsOrderOperateHistoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderOperateHistoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderOperateHistoryDo
	Group(cols ...field.Expr) IOmsOrderOperateHistoryDo
	Having(conds ...gen.Condition) IOmsOrderOperateHistoryDo
	Limit(limit int) IOmsOrderOperateHistoryDo
	Offset(offset int) IOmsOrderOperateHistoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderOperateHistoryDo
	Unscoped() IOmsOrderOperateHistoryDo
	Create(values ...*model.OmsOrderOperateHistory) error
	CreateInBatches(values []*model.OmsOrderOperateHistory, batchSize int) error
	Save(values ...*model.OmsOrderOperateHistory) error
	First() (*model.OmsOrderOperateHistory, error)
	Take() (*model.OmsOrderOperateHistory, error)
	Last() (*model.OmsOrderOperateHistory, error)
	Find() ([]*model.OmsOrderOperateHistory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderOperateHistory, err error)
	FindInBatches(result *[]*model.OmsOrderOperateHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OmsOrderOperateHistory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOmsOrderOperateHistoryDo
	Assign(attrs ...field.AssignExpr) IOmsOrderOperateHistoryDo
	Joins(fields ...field.RelationField) IOmsOrderOperateHistoryDo
	Preload(fields ...field.RelationField) IOmsOrderOperateHistoryDo
	FirstOrInit() (*model.OmsOrderOperateHistory, error)
	FirstOrCreate() (*model.OmsOrderOperateHistory, error)
	FindByPage(offset int, limit int) (result []*model.OmsOrderOperateHistory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOmsOrderOperateHistoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o omsOrderOperateHistoryDo) Debug() IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Debug())
}

func (o omsOrderOperateHistoryDo) WithContext(ctx context.Context) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o omsOrderOperateHistoryDo) ReadDB() IOmsOrderOperateHistoryDo {
	return o.Clauses(dbresolver.Read)
}

func (o omsOrderOperateHistoryDo) WriteDB() IOmsOrderOperateHistoryDo {
	return o.Clauses(dbresolver.Write)
}

func (o omsOrderOperateHistoryDo) Session(config *gorm.Session) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Session(config))
}

func (o omsOrderOperateHistoryDo) Clauses(conds ...clause.Expression) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o omsOrderOperateHistoryDo) Returning(value interface{}, columns ...string) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o omsOrderOperateHistoryDo) Not(conds ...gen.Condition) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o omsOrderOperateHistoryDo) Or(conds ...gen.Condition) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o omsOrderOperateHistoryDo) Select(conds ...field.Expr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o omsOrderOperateHistoryDo) Where(conds ...gen.Condition) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o omsOrderOperateHistoryDo) Order(conds ...field.Expr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o omsOrderOperateHistoryDo) Distinct(cols ...field.Expr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o omsOrderOperateHistoryDo) Omit(cols ...field.Expr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o omsOrderOperateHistoryDo) Join(table schema.Tabler, on ...field.Expr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o omsOrderOperateHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o omsOrderOperateHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o omsOrderOperateHistoryDo) Group(cols ...field.Expr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o omsOrderOperateHistoryDo) Having(conds ...gen.Condition) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o omsOrderOperateHistoryDo) Limit(limit int) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o omsOrderOperateHistoryDo) Offset(offset int) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o omsOrderOperateHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o omsOrderOperateHistoryDo) Unscoped() IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Unscoped())
}

func (o omsOrderOperateHistoryDo) Create(values ...*model.OmsOrderOperateHistory) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o omsOrderOperateHistoryDo) CreateInBatches(values []*model.OmsOrderOperateHistory, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o omsOrderOperateHistoryDo) Save(values ...*model.OmsOrderOperateHistory) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o omsOrderOperateHistoryDo) First() (*model.OmsOrderOperateHistory, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperateHistory), nil
	}
}

func (o omsOrderOperateHistoryDo) Take() (*model.OmsOrderOperateHistory, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperateHistory), nil
	}
}

func (o omsOrderOperateHistoryDo) Last() (*model.OmsOrderOperateHistory, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperateHistory), nil
	}
}

func (o omsOrderOperateHistoryDo) Find() ([]*model.OmsOrderOperateHistory, error) {
	result, err := o.DO.Find()
	return result.([]*model.OmsOrderOperateHistory), err
}

func (o omsOrderOperateHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OmsOrderOperateHistory, err error) {
	buf := make([]*model.OmsOrderOperateHistory, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o omsOrderOperateHistoryDo) FindInBatches(result *[]*model.OmsOrderOperateHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o omsOrderOperateHistoryDo) Attrs(attrs ...field.AssignExpr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o omsOrderOperateHistoryDo) Assign(attrs ...field.AssignExpr) IOmsOrderOperateHistoryDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o omsOrderOperateHistoryDo) Joins(fields ...field.RelationField) IOmsOrderOperateHistoryDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o omsOrderOperateHistoryDo) Preload(fields ...field.RelationField) IOmsOrderOperateHistoryDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o omsOrderOperateHistoryDo) FirstOrInit() (*model.OmsOrderOperateHistory, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperateHistory), nil
	}
}

func (o omsOrderOperateHistoryDo) FirstOrCreate() (*model.OmsOrderOperateHistory, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OmsOrderOperateHistory), nil
	}
}

func (o omsOrderOperateHistoryDo) FindByPage(offset int, limit int) (result []*model.OmsOrderOperateHistory, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o omsOrderOperateHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o omsOrderOperateHistoryDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o omsOrderOperateHistoryDo) Delete(models ...*model.OmsOrderOperateHistory) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *omsOrderOperateHistoryDo) withDO(do gen.Dao) *omsOrderOperateHistoryDo {
	o.DO = *do.(*gen.DO)
	return o
}
