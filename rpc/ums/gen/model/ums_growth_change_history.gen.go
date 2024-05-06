// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUmsGrowthChangeHistory = "ums_growth_change_history"

// UmsGrowthChangeHistory 成长值变化历史记录表
type UmsGrowthChangeHistory struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MemberID    int64     `gorm:"column:member_id;not null;comment:会员id" json:"member_id"`                    // 会员id
	CreateTime  time.Time `gorm:"column:create_time;not null;comment:创建时间" json:"create_time"`                // 创建时间
	ChangeType  int32     `gorm:"column:change_type;not null;comment:改变类型：0->增加；1->减少" json:"change_type"`    // 改变类型：0->增加；1->减少
	ChangeCount int32     `gorm:"column:change_count;not null;comment:积分改变数量" json:"change_count"`            // 积分改变数量
	OperateMan  string    `gorm:"column:operate_man;not null;comment:操作人员" json:"operate_man"`                // 操作人员
	OperateNote *string   `gorm:"column:operate_note;comment:操作备注" json:"operate_note"`                       // 操作备注
	SourceType  int32     `gorm:"column:source_type;not null;comment:积分来源：0->购物；1->管理员修改" json:"source_type"` // 积分来源：0->购物；1->管理员修改
}

// TableName UmsGrowthChangeHistory's table name
func (*UmsGrowthChangeHistory) TableName() string {
	return TableNameUmsGrowthChangeHistory
}
