package service

type LikeService interface {
	// LikeVideo 登录用户对视频的点赞和取消点赞操作
	LikeVideo(userID uint64, videoID uint64, actionType ActionType) error

	// LikeList 用户的所有点赞视频 ID
	LikeList(userID uint64) ([]int64, error)
}
