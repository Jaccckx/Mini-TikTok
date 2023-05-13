package dao

import (
	"io/ioutil"
	"math"
	"mime/multipart"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

type TableVideo struct {
	ID          int64 `gorm:"primaryKey;autoIncrement"`
	AuthorID    int64
	Title       string
	PlayURL     string
	CoverURL    string
	PublishTime time.Time
}

func (TableVideo) TableName() string {
	return "videos"
}

// breif：根据投稿时间，从数据库查询满足要求视频
// @para timeUnix time.Time：视频最晚发布时间
// @return tableVideos []TableVideo: 满足要求视频列表
// @return nextTime int64：视频列表中视频的最晚发布时间
func GetVideos(timeUnix time.Time) (tableVideos []TableVideo, nextTime int64) {
	result := Db.Model(&TableVideo{}).Where("publish_time < ?", timeUnix).Order("publish_time desc").Limit(30).Find(&tableVideos)
	logrus.Debug("Dao GetVideos:", tableVideos)

	if result.Error != nil {
		logrus.Error("Dao GetVideos:", result.Error)
		return
	}
	//处理nextTime
	if len(tableVideos) == 0 {
		nextTime = math.MaxInt64
	} else {
		nextTime = int64(tableVideos[len(tableVideos)-1].PublishTime.Unix())
	}
	return
}

func GetVideoListByUserID(userID int64) (tableVideos []TableVideo) {
	result := Db.Model(&TableVideo{}).Where("author_id = ?", userID).Order("publish_time desc").Find(&tableVideos)
	logrus.Debug("Dao GetVideoListByUserID:", tableVideos)

	if result.Error != nil {
		logrus.Error("Dao GetVideos:", result.Error)
		return
	}
	return
}

func GetLikeVideoListByUserID(userID int64) (tableVideos []TableVideo) {
	result := Db.Table("videos").Select("videos.*").
		Joins("JOIN likes ON videos.ID = likes.video_Id").
		Where("likes.user_Id = ? && likes.like = 1", userID).
		Find(&tableVideos)
	if result.Error != nil {
		logrus.Error("Dao GetVideos:", result.Error)
		return
	}
	return
}

// breif：根据用户ID，从数据库查询用户作品数
// @para userID int64：用户ID
// @return count int64: 用户作品数
// @return err error：报错
// BUG：是否支持滑到底，不显示操作？
func GetVideoCount(userID int64) (count int64, err error) {
	Db.Model(&TableVideo{}).Where("author_id = ?", userID).Count(&count)
	return
}

func GetFileToService(file *multipart.FileHeader) (path string, err error) {
	var dataStream multipart.File
	dataStream, err = file.Open()
	if err != nil {
		logrus.Info("publish流视频打开失败！")
		return
	}
	defer dataStream.Close()
	data := make([]byte, file.Size)
	if _, err = dataStream.Read(data); err != nil {
		logrus.Info("publish流视频读写文件失败！")
		return
	}

	path = "resources/upload/" + file.Filename
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		logrus.Println("publish流视频写文件失败！")
		return
	}
	return
}

func UploadFileToOss(path string, title string) (err error) {
	err = ossBucket.PutObjectFromFile(title, path)
	err = ossBucket.SetObjectACL(title, oss.ACLPublicReadWrite)
	if err != nil {
		logrus.Error("UploadFileToOss failed: ", err)
		return
	}
	return
}

func DeleteFileFromOss(title string) (err error) {
	err = ossBucket.DeleteObject(title)
	if err != nil {
		logrus.Error("UploadFileToOss failed: ", err)
		return
	}
	return
}

func GetUrlFromOss(fileName string) (PlayURL string, err error) {
	PlayURL = urlPrefix + fileName
	return
}

// 没有引入ffmeg
func InsertVideoRecordToDataBase(title string, userID int64, playUrl string, coverURL string) (videoID int64, err error) {
	video := &TableVideo{
		AuthorID:    userID,
		Title:       title,
		PlayURL:     playUrl,
		CoverURL:    coverURL,
		PublishTime: time.Now(),
	}
	result := Db.Model(&TableVideo{}).Create(&video)
	videoID = video.ID

	if result.Error != nil {
		logrus.Info("InsertVideoRecordToDataBase failed", result.Error)
		return -1, result.Error
	}
	return videoID, nil
}

func DeleteVideoRecordFromDataBase(videoID int64) (err error) {
	result := Db.Delete(&TableVideo{}, "id = ?", videoID)

	if result.Error != nil {
		logrus.Info("Delete video record failed:", result.Error)
		return result.Error
	}
	return nil
}

func ClearFileFromService(filePath string) {
	_, err := os.Stat(filePath)
	// 文件不存在
	if os.IsNotExist(err) {
		return
	}
	os.Remove(filePath)
	logrus.Debugln("ClearFileFromService succ!")
	return
}
