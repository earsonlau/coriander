package models

import "gorm.io/gorm"

type User struct {
	Model       gorm.Model `gorm:"embedded;embeddedPrefix:gorm_"`                       // 说明是嵌套字段，并添加嵌套前缀
	UserId      int64      `json:"user_id,string" gorm:"type:bigint;unique;index;not null;comment:用户id"`      // 指定在数据库中的字段类型，唯一，添加索引，非空
	Username    string     `gorm:"type:varchar(200);unique;not null;comment:用户名"` // 指定在数据库中的字段类型，唯一，添加索引，非空
	Password    string     `gorm:"type:varchar(200);not null;comment:密码"`               // 指定在数据库中的字段类型，非空
	Email       string     `gorm:"type:varchar(64);default:null;comment:邮箱"`            // 字段类型，默认为空
	PhoneNumber string     `gorm:"type:varchar(64);default:null;comment:电话号码"`          // 字段类型，默认为空
	Gender      int        `gorm:"type:tinyint;default:null;comment:性别[0-男 1-女]"`       // 字段类型，默认为空
	HeadImage   string     `gorm:"type:varchar(1000);default:null;comment:头像图片"`        // 字段类型，默认为空
}

