package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"log"
	"mini-tiktok/config"
)
import "gorm.io/gorm"

var Db *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/tiktok?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBUrl)
	fmt.Printf(dsn)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
}
