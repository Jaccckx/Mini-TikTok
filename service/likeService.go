package service

type LikeService interface {
	// LikeVideo 登录用户对视频的点赞和取消点赞操作
	LikeVideo(userID int64, videoID int64, actionType ActionType) error

	// LikeList 用户的所有点赞视频 ID
	LikeList(userID int64) ([]int64, error)

	// GetFavoriteCount 返回用户点赞数量
	GetFavoriteCount(userID int64) (int64, error)
}
