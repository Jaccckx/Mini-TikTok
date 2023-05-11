package service

import (
	"mini-tiktok/dao"
	"time"
)

type VideoServiceImpl struct {
}

func (v VideoServiceImpl) GetVideoCount(id int64) (int64, error) {
	return 0, nil
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
		if err != nil {
			videos[index].Author = *AuthorInfo
		}
	}
	return
}
