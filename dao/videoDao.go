package dao

import (
	"math"
	"time"

	"github.com/sirupsen/logrus"
)

type TableVideo struct {
	ID          int64 `gorm:"primaryKey"`
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
// BUG：是否支持滑到底，不显示操作？
func GetVideos(timeUnix time.Time) (tableVideos []TableVideo, nextTime int64) {
	Db.Where("publish_time < ?", timeUnix).Order("publish_time desc").Limit(30).Find(&tableVideos)
	logrus.Debug("Dao GetVideos:", tableVideos)

	//处理nextTime
	if len(tableVideos) == 0 {
		nextTime = math.MaxInt64
	} else {
		nextTime = int64(tableVideos[len(tableVideos)-1].PublishTime.Unix())
	}
	return
}
