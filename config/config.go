package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/db_go?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Product{})
	return db
}