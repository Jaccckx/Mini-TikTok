package dao

type Follows struct {
	ID          int64
	FollowingID int64
	FollowerID  int64
}

func (Follows) TableName() string {
	return "follow"
}

// InsertFollow 插入到数据库
func InsertFollow(followerID int64, followingID int64) error {
	result := Db.Create(&Follows{
		FollowerID:  followerID,
		FollowingID: followingID,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteFollow 删除数据
func DeleteFollow(followerID int64, followingID int64) error {
	result := Db.Where("following_id = ? AND follower_id = ?", followingID, followerID).Delete(&Follows{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetFollowList 返回关注的人
func GetFollowList(followerID int64) ([]int64, error) {
	var follows []*Follows
	result := Db.Where("follower_id =?", followerID).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}
	ids := make([]int64, len(follows))
	for i, v := range follows {
		ids[i] = v.FollowingID
	}
	return ids, nil
}

// GetFollowerList 返回粉丝列表
func GetFollowerList(followingID int64) ([]int64, error) {
	var follows []*Follows
	result := Db.Where("following_id =?", followingID).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}
	ids := make([]int64, len(follows))
	for i, v := range follows {
		ids[i] = v.FollowerID
	}
	return ids, nil
}

// GetFollowCount 返回关注数量
func GetFollowCount(followerID int64) (int64, error) {
	var count int64
	result := Db.Table("follow").Where("follower_id =?", followerID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// GetFollowerCount 返回粉丝数量
func GetFollowerCount(followingID int64) (int64, error) {
	var count int64
	result := Db.Table("follow").Where("following_id =?", followingID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// GetIsFollow 是否关注用户
func GetIsFollow(followingID int64, followsID int64) (bool, error) {
	var follow Follows
	result := Db.Where("following_id =? AND follower_id =?", followingID, followsID).Limit(1).Find(&follow)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected == 1, nil
}

// GetFriendList 返回朋友列表
func GetFriendList(userID int64) ([]int64, error) {
	var follow []*Follows
	// SELECT * FROM follows WHERE following_id in (SELECT follower_id FROM follows WHERE following_id = ?) AND follower_id = ?
	result := Db.Where("following_id in (?) AND follower_id = ?",
		Db.Table("follow").Where("following_id = ?", userID).Select("follower_id"),
		userID,
	).Find(&follow)

	if result.Error != nil {
		return nil, result.Error
	}

	ids := make([]int64, len(follow))
	for i, v := range follow {
		ids[i] = v.FollowingID
	}
	return ids, nil
}
