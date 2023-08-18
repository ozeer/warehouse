package models

import "gorm.io/gorm"

type RepoStarRecord struct {
	gorm.Model
	Uid int    `gorm:"column:uid;type:int(11);" json:"uid"` // uid star用户uid
	Rid string `gorm:"column:rid;type:int(11);" json:"rid"` // rid 仓库id
}

func (table *RepoStarRecord) TableName() string {
	return "repo_star_record"
}
