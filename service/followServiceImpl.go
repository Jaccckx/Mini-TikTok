package service

type FollowServiceImpl struct{}

func (FollowServiceImpl) FollowUser(userID int64, toUserID int64, actionType ActionType) error {
	//TODO implement me
	panic("implement me")
}

func (FollowServiceImpl) FollowUserList(userID int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (FollowServiceImpl) FollowerUserList(userID int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (FollowServiceImpl) FriendUserList(userID int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (FollowServiceImpl) GetFollowCount(userID int64) (int64, error) {
	//TODO implement me
	return 0, nil
}

func (FollowServiceImpl) GetFollowerCount(userID int64) (int64, error) {
	//TODO implement me
	return 0, nil
}

func (FollowServiceImpl) GetIsFollow(userID int64, followId int64) (bool, error) {
	//TODO implement me
	return false, nil
}
