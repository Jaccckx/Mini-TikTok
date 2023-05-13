package service

import (
	"mini-tiktok/dao"
)

type LikeServiceImpl struct{}

func (LikeServiceImpl) LikeVideo(userID int64, videoID int64, actionType string) error {
	return dao.ActionLikeVideo(userID, videoID, actionType)
}

func (LikeServiceImpl) GetFavoriteCountByUserID(userID int64) (int64, error) {
	return dao.GetFavoriteCountByUserID(userID)
}

func (LikeServiceImpl) IsFavoriteVideo(userID int64, videoID int64) (bool, error) {
	return dao.IsFavoriteVideo(userID, videoID)
}

func (LikeServiceImpl) GetFavoriteCountByVideoID(videoID int64) (int64, error) {
	return dao.GetFavoriteCountByVideoID(videoID)
}
