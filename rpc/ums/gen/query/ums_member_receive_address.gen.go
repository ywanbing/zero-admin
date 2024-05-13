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

func newUmsMemberReceiveAddress(db *gorm.DB, opts ...gen.DOOption) umsMemberReceiveAddress {
	_umsMemberReceiveAddress := umsMemberReceiveAddress{}

	_umsMemberReceiveAddress.umsMemberReceiveAddressDo.UseDB(db, opts...)
	_umsMemberReceiveAddress.umsMemberReceiveAddressDo.UseModel(&model.UmsMemberReceiveAddress{})

	tableName := _umsMemberReceiveAddress.umsMemberReceiveAddressDo.TableName()
	_umsMemberReceiveAddress.ALL = field.NewAsterisk(tableName)
	_umsMemberReceiveAddress.ID = field.NewInt64(tableName, "id")
	_umsMemberReceiveAddress.MemberID = field.NewInt64(tableName, "member_id")
	_umsMemberReceiveAddress.MemberName = field.NewString(tableName, "member_name")
	_umsMemberReceiveAddress.PhoneNumber = field.NewString(tableName, "phone_number")
	_umsMemberReceiveAddress.DefaultStatus = field.NewInt32(tableName, "default_status")
	_umsMemberReceiveAddress.PostCode = field.NewString(tableName, "post_code")
	_umsMemberReceiveAddress.Province = field.NewString(tableName, "province")
	_umsMemberReceiveAddress.City = field.NewString(tableName, "city")
	_umsMemberReceiveAddress.Region = field.NewString(tableName, "region")
	_umsMemberReceiveAddress.DetailAddress = field.NewString(tableName, "detail_address")
	_umsMemberReceiveAddress.CreateTime = field.NewTime(tableName, "create_time")
	_umsMemberReceiveAddress.UpdateTime = field.NewTime(tableName, "update_time")

	_umsMemberReceiveAddress.fillFieldMap()

	return _umsMemberReceiveAddress
}

// umsMemberReceiveAddress 会员收货地址表
type umsMemberReceiveAddress struct {
	umsMemberReceiveAddressDo umsMemberReceiveAddressDo

	ALL           field.Asterisk
	ID            field.Int64
	MemberID      field.Int64  // 会员id
	MemberName    field.String // 收货人名称
	PhoneNumber   field.String // 收货人电话
	DefaultStatus field.Int32  // 是否为默认
	PostCode      field.String // 邮政编码
	Province      field.String // 省份/直辖市
	City          field.String // 城市
	Region        field.String // 区
	DetailAddress field.String // 详细地址(街道)
	CreateTime    field.Time   // 创建时间
	UpdateTime    field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (u umsMemberReceiveAddress) Table(newTableName string) *umsMemberReceiveAddress {
	u.umsMemberReceiveAddressDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u umsMemberReceiveAddress) As(alias string) *umsMemberReceiveAddress {
	u.umsMemberReceiveAddressDo.DO = *(u.umsMemberReceiveAddressDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *umsMemberReceiveAddress) updateTableName(table string) *umsMemberReceiveAddress {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.MemberID = field.NewInt64(table, "member_id")
	u.MemberName = field.NewString(table, "member_name")
	u.PhoneNumber = field.NewString(table, "phone_number")
	u.DefaultStatus = field.NewInt32(table, "default_status")
	u.PostCode = field.NewString(table, "post_code")
	u.Province = field.NewString(table, "province")
	u.City = field.NewString(table, "city")
	u.Region = field.NewString(table, "region")
	u.DetailAddress = field.NewString(table, "detail_address")
	u.CreateTime = field.NewTime(table, "create_time")
	u.UpdateTime = field.NewTime(table, "update_time")

	u.fillFieldMap()

	return u
}

func (u *umsMemberReceiveAddress) WithContext(ctx context.Context) IUmsMemberReceiveAddressDo {
	return u.umsMemberReceiveAddressDo.WithContext(ctx)
}

func (u umsMemberReceiveAddress) TableName() string { return u.umsMemberReceiveAddressDo.TableName() }

func (u umsMemberReceiveAddress) Alias() string { return u.umsMemberReceiveAddressDo.Alias() }

func (u umsMemberReceiveAddress) Columns(cols ...field.Expr) gen.Columns {
	return u.umsMemberReceiveAddressDo.Columns(cols...)
}

func (u *umsMemberReceiveAddress) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *umsMemberReceiveAddress) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 12)
	u.fieldMap["id"] = u.ID
	u.fieldMap["member_id"] = u.MemberID
	u.fieldMap["member_name"] = u.MemberName
	u.fieldMap["phone_number"] = u.PhoneNumber
	u.fieldMap["default_status"] = u.DefaultStatus
	u.fieldMap["post_code"] = u.PostCode
	u.fieldMap["province"] = u.Province
	u.fieldMap["city"] = u.City
	u.fieldMap["region"] = u.Region
	u.fieldMap["detail_address"] = u.DetailAddress
	u.fieldMap["create_time"] = u.CreateTime
	u.fieldMap["update_time"] = u.UpdateTime
}

func (u umsMemberReceiveAddress) clone(db *gorm.DB) umsMemberReceiveAddress {
	u.umsMemberReceiveAddressDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u umsMemberReceiveAddress) replaceDB(db *gorm.DB) umsMemberReceiveAddress {
	u.umsMemberReceiveAddressDo.ReplaceDB(db)
	return u
}

type umsMemberReceiveAddressDo struct{ gen.DO }

type IUmsMemberReceiveAddressDo interface {
	gen.SubQuery
	Debug() IUmsMemberReceiveAddressDo
	WithContext(ctx context.Context) IUmsMemberReceiveAddressDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUmsMemberReceiveAddressDo
	WriteDB() IUmsMemberReceiveAddressDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUmsMemberReceiveAddressDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUmsMemberReceiveAddressDo
	Not(conds ...gen.Condition) IUmsMemberReceiveAddressDo
	Or(conds ...gen.Condition) IUmsMemberReceiveAddressDo
	Select(conds ...field.Expr) IUmsMemberReceiveAddressDo
	Where(conds ...gen.Condition) IUmsMemberReceiveAddressDo
	Order(conds ...field.Expr) IUmsMemberReceiveAddressDo
	Distinct(cols ...field.Expr) IUmsMemberReceiveAddressDo
	Omit(cols ...field.Expr) IUmsMemberReceiveAddressDo
	Join(table schema.Tabler, on ...field.Expr) IUmsMemberReceiveAddressDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberReceiveAddressDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberReceiveAddressDo
	Group(cols ...field.Expr) IUmsMemberReceiveAddressDo
	Having(conds ...gen.Condition) IUmsMemberReceiveAddressDo
	Limit(limit int) IUmsMemberReceiveAddressDo
	Offset(offset int) IUmsMemberReceiveAddressDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberReceiveAddressDo
	Unscoped() IUmsMemberReceiveAddressDo
	Create(values ...*model.UmsMemberReceiveAddress) error
	CreateInBatches(values []*model.UmsMemberReceiveAddress, batchSize int) error
	Save(values ...*model.UmsMemberReceiveAddress) error
	First() (*model.UmsMemberReceiveAddress, error)
	Take() (*model.UmsMemberReceiveAddress, error)
	Last() (*model.UmsMemberReceiveAddress, error)
	Find() ([]*model.UmsMemberReceiveAddress, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberReceiveAddress, err error)
	FindInBatches(result *[]*model.UmsMemberReceiveAddress, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UmsMemberReceiveAddress) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUmsMemberReceiveAddressDo
	Assign(attrs ...field.AssignExpr) IUmsMemberReceiveAddressDo
	Joins(fields ...field.RelationField) IUmsMemberReceiveAddressDo
	Preload(fields ...field.RelationField) IUmsMemberReceiveAddressDo
	FirstOrInit() (*model.UmsMemberReceiveAddress, error)
	FirstOrCreate() (*model.UmsMemberReceiveAddress, error)
	FindByPage(offset int, limit int) (result []*model.UmsMemberReceiveAddress, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUmsMemberReceiveAddressDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u umsMemberReceiveAddressDo) Debug() IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Debug())
}

func (u umsMemberReceiveAddressDo) WithContext(ctx context.Context) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u umsMemberReceiveAddressDo) ReadDB() IUmsMemberReceiveAddressDo {
	return u.Clauses(dbresolver.Read)
}

func (u umsMemberReceiveAddressDo) WriteDB() IUmsMemberReceiveAddressDo {
	return u.Clauses(dbresolver.Write)
}

func (u umsMemberReceiveAddressDo) Session(config *gorm.Session) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Session(config))
}

func (u umsMemberReceiveAddressDo) Clauses(conds ...clause.Expression) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u umsMemberReceiveAddressDo) Returning(value interface{}, columns ...string) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u umsMemberReceiveAddressDo) Not(conds ...gen.Condition) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u umsMemberReceiveAddressDo) Or(conds ...gen.Condition) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u umsMemberReceiveAddressDo) Select(conds ...field.Expr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u umsMemberReceiveAddressDo) Where(conds ...gen.Condition) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u umsMemberReceiveAddressDo) Order(conds ...field.Expr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u umsMemberReceiveAddressDo) Distinct(cols ...field.Expr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u umsMemberReceiveAddressDo) Omit(cols ...field.Expr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u umsMemberReceiveAddressDo) Join(table schema.Tabler, on ...field.Expr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u umsMemberReceiveAddressDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u umsMemberReceiveAddressDo) RightJoin(table schema.Tabler, on ...field.Expr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u umsMemberReceiveAddressDo) Group(cols ...field.Expr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u umsMemberReceiveAddressDo) Having(conds ...gen.Condition) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u umsMemberReceiveAddressDo) Limit(limit int) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u umsMemberReceiveAddressDo) Offset(offset int) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u umsMemberReceiveAddressDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u umsMemberReceiveAddressDo) Unscoped() IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Unscoped())
}

func (u umsMemberReceiveAddressDo) Create(values ...*model.UmsMemberReceiveAddress) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u umsMemberReceiveAddressDo) CreateInBatches(values []*model.UmsMemberReceiveAddress, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u umsMemberReceiveAddressDo) Save(values ...*model.UmsMemberReceiveAddress) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u umsMemberReceiveAddressDo) First() (*model.UmsMemberReceiveAddress, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberReceiveAddress), nil
	}
}

func (u umsMemberReceiveAddressDo) Take() (*model.UmsMemberReceiveAddress, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberReceiveAddress), nil
	}
}

func (u umsMemberReceiveAddressDo) Last() (*model.UmsMemberReceiveAddress, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberReceiveAddress), nil
	}
}

func (u umsMemberReceiveAddressDo) Find() ([]*model.UmsMemberReceiveAddress, error) {
	result, err := u.DO.Find()
	return result.([]*model.UmsMemberReceiveAddress), err
}

func (u umsMemberReceiveAddressDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UmsMemberReceiveAddress, err error) {
	buf := make([]*model.UmsMemberReceiveAddress, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u umsMemberReceiveAddressDo) FindInBatches(result *[]*model.UmsMemberReceiveAddress, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u umsMemberReceiveAddressDo) Attrs(attrs ...field.AssignExpr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u umsMemberReceiveAddressDo) Assign(attrs ...field.AssignExpr) IUmsMemberReceiveAddressDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u umsMemberReceiveAddressDo) Joins(fields ...field.RelationField) IUmsMemberReceiveAddressDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u umsMemberReceiveAddressDo) Preload(fields ...field.RelationField) IUmsMemberReceiveAddressDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u umsMemberReceiveAddressDo) FirstOrInit() (*model.UmsMemberReceiveAddress, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberReceiveAddress), nil
	}
}

func (u umsMemberReceiveAddressDo) FirstOrCreate() (*model.UmsMemberReceiveAddress, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UmsMemberReceiveAddress), nil
	}
}

func (u umsMemberReceiveAddressDo) FindByPage(offset int, limit int) (result []*model.UmsMemberReceiveAddress, count int64, err error) {
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

func (u umsMemberReceiveAddressDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u umsMemberReceiveAddressDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u umsMemberReceiveAddressDo) Delete(models ...*model.UmsMemberReceiveAddress) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *umsMemberReceiveAddressDo) withDO(do gen.Dao) *umsMemberReceiveAddressDo {
	u.DO = *do.(*gen.DO)
	return u
}