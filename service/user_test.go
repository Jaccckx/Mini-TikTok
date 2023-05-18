package service

import (
	"github.com/go-playground/assert/v2"
	"mini-tiktok/dao"
	"mini-tiktok/middleware/redis"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	dao.Init()
	redis.Init()

	usi := UserServiceImpl{}
	info, err := usi.getUserInfo(60)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, int64(60), info.ID)
}

func TestInsertUser(t *testing.T) {
	dao.Init()
	redis.Init()

	usi := UserServiceImpl{}
	id, err := usi.InsertUser("test_user_12", "test_pwd_1")
	if err != nil {
		t.Error(err)
	}

	info, err := usi.GetUserInfoById(id, id)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, id, info.ID)
}
