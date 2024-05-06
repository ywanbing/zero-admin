// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsProductCategoryAttributeRelation = "pms_product_category_attribute_relation"

// PmsProductCategoryAttributeRelation 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type PmsProductCategoryAttributeRelation struct {
	ID                 int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProductCategoryID  int64 `gorm:"column:product_category_id;not null" json:"product_category_id"`
	ProductAttributeID int64 `gorm:"column:product_attribute_id;not null" json:"product_attribute_id"`
}

// TableName PmsProductCategoryAttributeRelation's table name
func (*PmsProductCategoryAttributeRelation) TableName() string {
	return TableNamePmsProductCategoryAttributeRelation
}
