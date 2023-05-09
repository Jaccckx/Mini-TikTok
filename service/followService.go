package service

type FollowService interface {
	// FollowUser 关注操作
	FollowUser(userID uint64, toUserID uint64, actionType ActionType) error

	// FollowUserList 返回关注列表用户 ID
	FollowUserList(userID uint64) ([]uint64, error)

	// FollowerUserList 返回粉丝列表用户 ID
	FollowerUserList(userID uint64) ([]uint64, error)

	// FriendUserList 返回朋友列表用户 ID
	FriendUserList(userID uint64) ([]uint64, error)
}
