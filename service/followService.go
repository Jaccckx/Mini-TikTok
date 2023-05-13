package service

type FollowService interface {
	// FollowUser 关注
	FollowUser(userID int64, followingID int64) error

	// UnFollowUser 取消关注
	UnFollowUser(userID int64, followingID int64) error

	// GetFollowList 返回关注 userID 列表
	GetFollowList(userID int64) ([]int64, error)

	// GetFollowerList 返回粉丝 userID 列表
	GetFollowerList(userID int64) ([]int64, error)

	// GetFriendList 返回朋友列表用户 ID
	GetFriendList(userID int64) ([]int64, error)

	// GetFollowCount 返回 UserID 的关注数量
	GetFollowCount(userID int64) (int64, error)

	// GetFollowerCount 返回 UserID 的粉丝数量
	GetFollowerCount(userID int64) (int64, error)

	// GetIsFollow 判断用户是否关注
	GetIsFollow(userID int64, followingID int64) (bool, error)
}
