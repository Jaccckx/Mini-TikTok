package service

type VideoService interface {
	// GetVideoCount 返回用户的投稿数量
	GetVideoCount(id int64) (int64, error)
}
