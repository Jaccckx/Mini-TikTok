package dao

import (
	"fmt"
	"log"
	"mini-tiktok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/tiktok?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBUrl)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
}
