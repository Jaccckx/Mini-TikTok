package service

import "mini-tiktok/dao"

type FollowServiceImpl struct{}

func (FollowServiceImpl) FollowUser(userID int64, followingID int64) error {
	return dao.InsertFollow(userID, followingID)
}

func (FollowServiceImpl) UnFollowUser(userID int64, followingID int64) error {
	return dao.DeleteFollow(userID, followingID)
}

func (FollowServiceImpl) GetFollowList(userID int64) ([]int64, error) {
	return dao.GetFollowList(userID)
}

func (FollowServiceImpl) GetFollowerList(userID int64) ([]int64, error) {
	return dao.GetFollowerList(userID)
}

func (FollowServiceImpl) GetFriendList(userID int64) ([]int64, error) {
	return dao.GetFriendList(userID)
}

func (FollowServiceImpl) GetFollowCount(userID int64) (int64, error) {
	return dao.GetFollowCount(userID)
}

func (FollowServiceImpl) GetFollowerCount(userID int64) (int64, error) {
	return dao.GetFollowerCount(userID)
}

func (FollowServiceImpl) GetIsFollow(userID int64, followId int64) (bool, error) {
	return dao.GetIsFollow(userID, followId)
}
