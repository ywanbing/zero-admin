// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsAlbum = "pms_album"

// PmsAlbum 相册表
type PmsAlbum struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string `gorm:"column:name;not null" json:"name"`
	CoverPic    string `gorm:"column:cover_pic;not null" json:"cover_pic"`
	PicCount    int32  `gorm:"column:pic_count;not null" json:"pic_count"`
	Sort        int32  `gorm:"column:sort;not null" json:"sort"`
	Description string `gorm:"column:description;not null" json:"description"`
}

// TableName PmsAlbum's table name
func (*PmsAlbum) TableName() string {
	return TableNamePmsAlbum
}