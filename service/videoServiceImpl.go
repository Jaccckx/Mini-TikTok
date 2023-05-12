package service

import (
	"mime/multipart"
	"mini-tiktok/dao"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type VideoServiceImpl struct {
}

func (v VideoServiceImpl) GetVideoCount(userID int64) (int64, error) {
	return dao.GetVideoCount(userID)
}

func (v VideoServiceImpl) fillVideos(tableVideos []dao.TableVideo, userID int64) (videos []Video) {
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
		if err == nil {
			videos[index].IsFavorite = videoisFavorite
		}
		videoFavoriteCount, err := LikeServiceImpl{}.GetFavoriteCountByVideoID(videoID)
		if err == nil {
			videos[index].FavoriteCount = videoFavoriteCount
		}
		videoCommentCount, err := (&CommentServiceImpl{}).GetCommentCountByVideoID(videoID)
		if err == nil {
			videos[index].CommentCount = videoCommentCount
		}
		AuthorInfo, err := (&UserServiceImpl{}).GetUserInfoById(authorID, userID)
		if err == nil {
			videos[index].Author = *AuthorInfo
		} else {
			logrus.Errorf("fillVideos Get AuthorInfo failed: %v", err)
		}
	}

	return
}

func (v VideoServiceImpl) GetVideos(timeUnix time.Time, userID int64) (videos []Video, nextTime int64) {
	//获取视频列表
	var tableVideos []dao.TableVideo
	tableVideos, nextTime = dao.GetVideos(timeUnix)
	videos = v.fillVideos(tableVideos, userID)
	return
}

func (v VideoServiceImpl) GetVideosList(userID int64) (videos []Video) {
	//获取视频列表
	var tableVideos []dao.TableVideo
	tableVideos = dao.GetVideoListByUserID(userID)
	videos = v.fillVideos(tableVideos, userID)
	return
}

func (v VideoServiceImpl) PublishVideo(file *multipart.FileHeader, userID int64, title string) error {
	path, err := dao.GetFileToService(file)
	if err != nil {
		logrus.Error("[VideoServiceImpl-PublishVideo] GetFileToService failed: ", err)
		return err
	}

	cmd := exec.Command("ffmpeg", "-ss", "00:00:01", "-i", path, "-frames:v", "1", "./resources/upload/"+title+".jpg")
	if err := cmd.Run(); err != nil {
		logrus.Error("[VideoServiceImpl-GetPic] GetPic failed")
	}

	err = dao.UploadFileToOss("./resources/upload/"+title+".jpg", title+".jpg")
	if err != nil {
		logrus.Error("[VideoServiceImpl-PublishVideo] UploadFileToOss failed: ", err)
		return err
	}
	err = dao.UploadFileToOss(path, title+filepath.Ext(file.Filename))
	if err != nil {
		dao.DeleteFileFromOss(title + ".jpg")
		logrus.Error("[VideoServiceImpl-PublishVideo] UploadFileToOss failed: ", err)
		return err
	}

	CoverUrl, err := dao.GetUrlFromOss(title + ".jpg")
	playUrl, err := dao.GetUrlFromOss(title + filepath.Ext(file.Filename))
	if err != nil {
		logrus.Error("[VideoServiceImpl-PublishVideo] GetUrlFromOss failed: ", err)
		dao.DeleteFileFromOss(title + filepath.Ext(file.Filename))
		dao.DeleteFileFromOss(title + ".jpg")
		return err
	}

	_, err = dao.InsertVideoRecordToDataBase(title, userID, playUrl, CoverUrl)
	if err != nil {
		logrus.Error("[VideoServiceImpl-PublishVideo] InsertVideoRecordToDataBase failed: ", err)
		dao.DeleteFileFromOss(title + filepath.Ext(file.Filename))
		dao.DeleteFileFromOss(title + ".jpg")
		return err
	}

	dao.ClearFileFromService(path)
	return nil
}
