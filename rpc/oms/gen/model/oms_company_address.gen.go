// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOmsCompanyAddress = "oms_company_address"

// OmsCompanyAddress 公司收发货地址表
type OmsCompanyAddress struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AddressName   string `gorm:"column:address_name;not null;comment:地址名称" json:"address_name"`                   // 地址名称
	SendStatus    int32  `gorm:"column:send_status;not null;comment:默认发货地址：0->否；1->是" json:"send_status"`         // 默认发货地址：0->否；1->是
	ReceiveStatus int32  `gorm:"column:receive_status;not null;comment:是否默认收货地址：0->否；1->是" json:"receive_status"` // 是否默认收货地址：0->否；1->是
	Name          string `gorm:"column:name;not null;comment:收发货人姓名" json:"name"`                                 // 收发货人姓名
	Phone         string `gorm:"column:phone;not null;comment:收货人电话" json:"phone"`                                // 收货人电话
	Province      string `gorm:"column:province;not null;comment:省/直辖市" json:"province"`                          // 省/直辖市
	City          string `gorm:"column:city;not null;comment:市" json:"city"`                                      // 市
	Region        string `gorm:"column:region;not null;comment:区" json:"region"`                                  // 区
	DetailAddress string `gorm:"column:detail_address;not null;comment:详细地址" json:"detail_address"`               // 详细地址
}

// TableName OmsCompanyAddress's table name
func (*OmsCompanyAddress) TableName() string {
	return TableNameOmsCompanyAddress
}
