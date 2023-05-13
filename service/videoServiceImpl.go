package service

import (
	"mime/multipart"
	"mini-tiktok/dao"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type VideoServiceImpl struct {
}

func (v VideoServiceImpl) GetVideoCount(userID int64) (int64, error) {
	return dao.GetVideoCount(userID)
}

func (v VideoServiceImpl) GetVideos(timeUnix time.Time, userID int64) (videos []Video, nextTime int64) {
	//获取视频列表
	var tableVideos []dao.TableVideo
	tableVideos, nextTime = dao.GetVideos(timeUnix)
	for _, value := range tableVideos {
		video := Video{
			Author:   User{},
			CoverURL: value.CoverURL,
			ID:       value.ID,
			PlayURL:  value.PlayURL,
			Title:    value.Title,
		}
		video.Author.ID = value.AuthorID
		videos = append(videos, video)
	}

	for index := range videos {
		videoID := videos[index].ID
		authorID := videos[index].Author.ID
		videoisFavorite, err := LikeServiceImpl{}.IsFavoriteVideo(userID, videoID)
		if err != nil {
			videos[index].IsFavorite = videoisFavorite
		}
		videoFavoriteCount, err := LikeServiceImpl{}.GetFavoriteCountByVideoID(videoID)
		if err != nil {
			videos[index].FavoriteCount = videoFavoriteCount
		}
		videoCommentCount, err := (&CommentServiceImpl{}).GetCommentCountByVideoID(videoID)
		if err != nil {
			videos[index].CommentCount = videoCommentCount
		}
		AuthorInfo, err := (&UserServiceImpl{}).GetUserInfoById(authorID, userID)
		if err == nil {
			videos[index].Author = *AuthorInfo
		} else {
			logrus.Error("[VideoServiceImpl-GetVideos] GetUserInfoById error: ", err)
		}
	}
	return
}

// PublishVideo BUG：上传的视频无法播放！
func (v VideoServiceImpl) PublishVideo(file *multipart.FileHeader, userID int64, title string) error {
	path, err := dao.GetFileToService(file)
	if err != nil {
		logrus.Error("[VideoServiceImpl-PublishVideo] GetFileToService failed: ", err)
		return err
	}

	err = dao.UploadFileToOss(path, title+filepath.Ext(file.Filename))
	if err != nil {
		logrus.Error("[VideoServiceImpl-PublishVideo] UploadFileToOss failed: ", err)
		return err
	}

	playUrl, err := dao.GetUrlFromOss(title + filepath.Ext(file.Filename))
	if err != nil {
		logrus.Error("[VideoServiceImpl-PublishVideo] GetUrlFromOss failed: ", err)
		return err
	}

	err = dao.InsertVideoRecordToDataBase(title, userID, playUrl, playUrl)
	if err != nil {
		logrus.Error("[VideoServiceImpl-PublishVideo] InsertVideoRecordToDataBase failed: ", err)
		dao.DeleteFileFromOss(title + filepath.Ext(file.Filename))
		return err
	}

	dao.ClearFileFromService(path)
	return nil
}
