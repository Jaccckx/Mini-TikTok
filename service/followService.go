package service

type FollowService interface {
	// FollowUser 关注操作
	FollowUser(userID int64, toUserID int64, actionType ActionType) error

	// FollowUserList 返回关注列表用户 ID
	FollowUserList(userID int64) ([]int64, error)

	// FollowerUserList 返回粉丝列表用户 ID
	FollowerUserList(userID int64) ([]int64, error)

	// FriendUserList 返回朋友列表用户 ID
	FriendUserList(userID int64) ([]int64, error)

	// GetFollowCount 返回 UserID 的粉丝数量
	GetFollowCount(userID int64) (int64, error)

	// GetFollowerCount 返回 UserID 的关注数量
	GetFollowerCount(userID int64) (int64, error)

	GetIsFollow(userID int64, followId int64) (bool, error)
}
