package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`  // 唯一标识
	Uid      string `gorm:"column:uid;type:varchar(36);" json:"uid"`            // 用户uid
	UserName string `gorm:"column:username;type:varchar(255);" json:"username"` // username 用户名
	Password string `gorm:"column:password;type:varchar(36);" json:"password"`  // password 密码
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
