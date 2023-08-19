package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(dns string) error {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalf("数据库连接错误: %s", err.Error())
		return err
	}

	err = db.AutoMigrate(&UserBasic{}, &RepoBasic{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}
