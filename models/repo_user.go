package models

import "gorm.io/gorm"

type RepoUser struct {
	gorm.Model
	Uid  int    `gorm:"column:uid;type:int(11);" json:"uid"`      // uid 用户uid
	Rid  string `gorm:"column:rid;type:int(11);" json:"rid"`      // rid 仓库id
	Type string `gorm:"column:type;type:tinyint(1);" json:"type"` // type 类型{1:所有者2:被授权者}
}

func (table *RepoUser) TableName() string {
	return "repo_user"
}
