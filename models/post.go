package models

import "gorm.io/gorm"

type Post struct {
	Model        gorm.Model `gorm:"embedded;embeddedPrefix:gorm_"`                         // 说明是嵌套字段，并添加嵌套前缀
	PostId      int64      `gorm:"type:bigint;unique;index;not null;comment:帖子id"`      // 指定在数据库中的字段类型，唯一，添加索引，非空
	Title    string     `gorm:"type:varchar(128);not null;comment:帖子标题"` // 指定在数据库中的字段类型，唯一，添加索引，非空
	Content    string     `gorm:"type:varchar(8192);not null;comment:帖子内容"`               // 指定在数据库中的字段类型，非空
	AuthorId       int64     `gorm:"type:bigint;default:null;comment:作者id"`            // 字段类型，默认为空
	CategoryId int64     `gorm:"type:bigint;default:null;comment:分类id"`          // 字段类型，默认为空
	Status      int        `gorm:"type:tinyint;default:null;comment:帖子状态[0-不存在 1-存在]"`       // 字段类型，默认为空
}

//帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Post //嵌入帖子结构体
	*Category `json:"category"` //嵌入类别信息
}