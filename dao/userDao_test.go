package dao

import (
	"log"
	"strconv"
	"testing"
)

func buildUser(num int) []*User {
	var users []*User
	for i := 0; i < num; i++ {
		users = append(users, &User{
			Name:            "test_" + strconv.Itoa(i),
			Password:        "123456",
			Avatar:          "123456",
			BackgroundImage: "123456",
			Signature:       "123456",
		})
	}
	return users
}

func cleanTestDB() {
	Db.Where("name LIKE ?", "test%").Delete(&User{})
}

func TestInsertTest(t *testing.T) {
	Init()
	users := buildUser(10)
	result := Db.Create(&users)
	if result.Error != nil {
		log.Panic(result.Error)
	}
	if result.RowsAffected != 10 {
		log.Panicf("insert num error: %v", result.RowsAffected)
	}
	cleanTestDB()
}
