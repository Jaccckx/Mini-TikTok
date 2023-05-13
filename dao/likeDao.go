package dao

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Like 表的结构。
type Like struct {
	UserId  int64 `gorm:"primaryKey;"`
	VideoId int64 `gorm:"primaryKey"`
	Like    int8  `gorm:"type:tinyint;not null"`
}

// TableName 修改表名映射
func (Like) TableName() string {
	return "likes"
}

func IsFavoriteVideo(userID int64, videoID int64) (bool, error) {
	var like Like
	// 查询是否存在对应记录
	err := Db.Where("user_Id = ? AND video_Id = ?", userID, videoID).First(&like).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询出错
		return false, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// 未查询到记录，新建一条
		like = Like{
			UserId:  userID,
			VideoId: videoID,
			Like:    0,
		}
		err := Db.Create(&like).Error
		if err != nil {
			return false, err
		}
	}
	return like.Like == 1, nil
}

func ActionLikeVideo(userID int64, videoID int64, actionType string) error {
	IsFavoriteVideo(userID, videoID)

	var like Like
	result := Db.Where("user_Id = ? AND video_Id = ?", userID, videoID).First(&like)
	if result.Error != nil {
		return result.Error
	}
	if actionType == "1" {
		like.Like = 1
	} else {
		like.Like = 0
	}

	like.UserId = userID
	like.VideoId = videoID
	result = Db.Save(&like)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetFavoriteCountByUserID(userID int64) (count int64, err error) {
	result := Db.Model(&Like{}).Where("user_Id = ? AND `like` = ?", userID, 1).Count(&count)
	if result.Error != nil {
		err = result.Error
		logrus.Info("GetFavoriteCountByUserID Failed", err)
		return
	}
	return
}

func GetFavoriteCountByVideoID(videoID int64) (count int64, err error) {
	result := Db.Model(&Like{}).Where("video_Id = ? AND `like` = ?", videoID, 1).Count(&count)
	if result.Error != nil {
		err = result.Error
		logrus.Info("GetFavoriteCountByVideoID Failed", err)
		return
	}
	return
}
