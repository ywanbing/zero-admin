// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUmsMemberProductCollection = "ums_member_product_collection"

// UmsMemberProductCollection 用户收藏的商品
type UmsMemberProductCollection struct {
	ID              int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MemberID        int64      `gorm:"column:member_id;not null;comment:会员id" json:"member_id"`                      // 会员id
	MemberNickName  string     `gorm:"column:member_nick_name;not null;comment:会员姓名" json:"member_nick_name"`        // 会员姓名
	MemberIcon      string     `gorm:"column:member_icon;not null;comment:会员头像" json:"member_icon"`                  // 会员头像
	ProductID       int64      `gorm:"column:product_id;not null;comment:商品id" json:"product_id"`                    // 商品id
	ProductName     string     `gorm:"column:product_name;not null;comment:商品名称" json:"product_name"`                // 商品名称
	ProductPic      string     `gorm:"column:product_pic;not null;comment:商品图片" json:"product_pic"`                  // 商品图片
	ProductSubTitle *string    `gorm:"column:product_sub_title;comment:商品标题" json:"product_sub_title"`               // 商品标题
	ProductPrice    float64    `gorm:"column:product_price;not null;comment:商品价格" json:"product_price"`              // 商品价格
	CreateTime      *time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:收藏时间" json:"create_time"` // 收藏时间
}

// TableName UmsMemberProductCollection's table name
func (*UmsMemberProductCollection) TableName() string {
	return TableNameUmsMemberProductCollection
}
