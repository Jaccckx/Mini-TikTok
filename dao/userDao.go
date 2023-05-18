package dao

import "github.com/sirupsen/logrus"

type User struct {
	ID              int64
	Name            string
	Password        string
	Avatar          string
	BackgroundImage string
	Signature       string
}

// InsertUser 插入到数据库, 返回插入的ID, 若失败返回 0
func InsertUser(user *User) (int64, error) {
	// 为了避免重复插入导致存在多个相同名字的用户，name 字段应该使用唯一索引
	result := Db.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return user.ID, nil
}

// GetUserIdByName 返回对应 name 的 user id, 若不存在返回 0
func GetUserIdByName(name string) (int64, error) {
	var user User
	result := Db.Where("name =?", name).Limit(1).Find(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return user.ID, nil
}

// GetUserIdByPassword 返回对应的 user id, 若不存在返回 0
func GetUserIdByPassword(name string, password string) (int64, error) {
	var user User
	result := Db.Where("name =? AND password =?", name, password).Limit(1).Find(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return user.ID, nil
}

func GetUserInfoByID(id int64) (*User, error) {
	var user User
	result := Db.First(&user, id)
	if result.Error != nil {
		logrus.Info("GetUserInfoByID failed: ", result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func GetUserInfoByName(name string) (*User, error) {
	var user User
	result := Db.Where("name =?", name).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
