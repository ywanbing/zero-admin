// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOmsOrderSetting = "oms_order_setting"

// OmsOrderSetting 订单设置表
type OmsOrderSetting struct {
	ID                  int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	FlashOrderOvertime  int32 `gorm:"column:flash_order_overtime;not null;comment:秒杀订单超时关闭时间(分)" json:"flash_order_overtime"` // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int32 `gorm:"column:normal_order_overtime;not null;comment:正常订单超时时间(分)" json:"normal_order_overtime"` // 正常订单超时时间(分)
	ConfirmOvertime     int32 `gorm:"column:confirm_overtime;not null;comment:发货后自动确认收货时间（天）" json:"confirm_overtime"`        // 发货后自动确认收货时间（天）
	FinishOvertime      int32 `gorm:"column:finish_overtime;not null;comment:自动完成交易时间，不能申请售后（天）" json:"finish_overtime"`      // 自动完成交易时间，不能申请售后（天）
	CommentOvertime     int32 `gorm:"column:comment_overtime;not null;comment:订单完成后自动好评时间（天）" json:"comment_overtime"`        // 订单完成后自动好评时间（天）
}

// TableName OmsOrderSetting's table name
func (*OmsOrderSetting) TableName() string {
	return TableNameOmsOrderSetting
}
