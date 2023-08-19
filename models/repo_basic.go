package models

import "gorm.io/gorm"

type RepoBasic struct {
	gorm.Model
	Uid      int    `gorm:"column:uid;type:int(11);" json:"uid"`               // uid 用户uid
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // identity 仓库唯一标识
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`        // name 仓库名
	Path     string `gorm:"column:path;type:varchar(255);" json:"path"`        // path 仓库路径
	Desc     string `gorm:"column:desc;type:varchar(255);" json:"desc"`        // desc 仓库描述
	Star     int32  `gorm:"column:star;type:int(11);default:0;" json:"star"`   // star 仓库星数
}

func (table *RepoBasic) TableName() string {
	return "repo_basic"
}
