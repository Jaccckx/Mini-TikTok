package service

import (
	"github.com/go-playground/assert/v2"
	"mini-tiktok/dao"
	"testing"
)

func TestFollowBasic(t *testing.T) {
	dao.Init()
	fsi := FollowServiceImpl{}

	err := fsi.FollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}

	list, err := fsi.GetFollowList(1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(list), 1)
	assert.Equal(t, list[0], int64(2))

	err = fsi.UnFollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}

	list, err = fsi.GetFollowerList(1)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(list), 0)

	err = fsi.UnFollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}
}

func TestFollowList(t *testing.T) {
	dao.Init()
	fsi := FollowServiceImpl{}

	err := fsi.FollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}

	err = fsi.FollowUser(1, 3)
	if err != nil {
		t.Error(err)
	}

	list, err := fsi.GetFollowerList(1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(list), 2)

	err = fsi.UnFollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}

	list, err = fsi.GetFollowerList(1)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(list), 1)

	err = fsi.UnFollowUser(1, 3)
	if err != nil {
		t.Error(err)
	}
}

func TestFollowCount(t *testing.T) {
	dao.Init()
	fsi := FollowServiceImpl{}

	err := fsi.FollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}

	err = fsi.FollowUser(1, 3)
	if err != nil {
		t.Error(err)
	}

	err = fsi.FollowUser(1, 4)
	if err != nil {
		t.Error(err)
	}

	c, err := fsi.GetFollowCount(1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, c, int64(3))

	c, err = fsi.GetFollowerCount(2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, c, int64(1))

	err = fsi.UnFollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}
	err = fsi.UnFollowUser(1, 3)
	if err != nil {
		t.Error(err)
	}
	err = fsi.UnFollowUser(1, 4)
	if err != nil {
		t.Error(err)
	}

	c, err = fsi.GetFollowCount(2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, c, int64(0))

	c, err = fsi.GetFollowerCount(1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, c, int64(0))
}

func TestFollowFriend(t *testing.T) {
	dao.Init()
	fsi := FollowServiceImpl{}

	err := fsi.FollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}

	err = fsi.FollowUser(2, 1)
	if err != nil {
		t.Error(err)
	}

	list1, err := fsi.GetFriendList(1)
	if err != nil {
		t.Error(err)
	}

	list2, err := fsi.GetFriendList(2)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(list1), 1)
	assert.Equal(t, len(list2), 1)
	assert.Equal(t, list1[0], int64(2))
	assert.Equal(t, list2[0], int64(1))

	err = fsi.UnFollowUser(1, 2)
	if err != nil {
		t.Error(err)
	}

	err = fsi.UnFollowUser(2, 1)
	if err != nil {
		t.Error(err)
	}
}
