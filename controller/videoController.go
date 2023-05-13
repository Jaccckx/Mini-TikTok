package controller

import (
	"mini-tiktok/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Feed breif：从gin.Context请求里提取lastest_time参数和token参数，并返回视频列表
// @para c *gin.Context：HTTP请求上下文
//
// 函数内容：
// 1.提取lastest_time参数，并调用GetTime函数处理格式
// 2.提取用户ID
// 3.根据lastest_time参数，返回视频列表。
// 4.如果视频列表为空，statusCode设为http.StatusNotFound，StatusMsg设为GetVideoFail
// BUG:点赞模块需要和用户模块联调
func Feed(c *gin.Context) {
	timeStr := c.Query("lastest_time")
	timeUnix := GetTime(timeStr)

	currUserID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)

	statusCode := int64(0)
	statusMsg := &getVideoSucc
	videoServiceTemp := service.VideoServiceImpl{}
	videoList, nextTime := videoServiceTemp.GetVideos(timeUnix, currUserID)
	if len(videoList) == 0 {
		statusCode = 1
		statusMsg = &emptyVideoErr
		logrus.Println("feed stream Get Empty VideoList")
	}

	c.JSON(200, FeedResponse{
		NextTime:   &nextTime,
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
		VideoList:  videoList,
	})

}

// GetTime breif：处理latest_time参数，使其由string格式变成time.Time格式
// @para timeStr string：lastest_time参数string格式
// @return timeUnix time.Time: lastest_time参数time.Time格式
//
// 函数内容：
// 1.如果传入timeStr为空，说明lastest_time传入参数为空，将time.Unix设为当前时间即可
// 2.如果解析timeStr失败，将time.Unix设为当前时间即可
// 3.解析timeStr成功，将传入timeStr转为time.Time格式
func GetTime(timeStr string) (timeUnix time.Time) {
	if len(timeStr) == 0 {
		timeUnix = time.Now()
		return
	}

	timeInt, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		timeUnix = time.Now()
		logrus.Info("feed stream ParseTime failed")
		return
	}
	timeUnix = time.Unix(timeInt, 0)
	return
}

func Publish(c *gin.Context) {
	fileHeader, err := c.FormFile("data")
	if err != nil {
		logrus.Info("[videoController-Publish] 获取视频流失败:%v", err)
		c.JSON(http.StatusOK, PublishListResponse{
			StatusCode: 1,
			StatusMsg:  "获取视频流失败",
		})
		return
	}

	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	logrus.Debugln("[videoController-Publish] Upload-userID: ", userId)

	title := c.PostForm("title")
	logrus.Debugln("[videoController-Publish] Upload-videoTitle: ", title)

	publishService := service.VideoServiceImpl{}
	err = publishService.PublishVideo(fileHeader, userId, title)

	statusCode := int64(0)
	statusMsg := "Upload Successfully"
	if err != nil {
		statusCode = 1
		statusMsg = "Upload Failed"
	}

	c.JSON(200, PublishActionResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	})
}

func Publishlist(c *gin.Context) {
	currUserID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	videoServiceTemp := service.VideoServiceImpl{}
	videoList := videoServiceTemp.GetPublishVideosList(currUserID)
	c.JSON(200, PublishListResponse{
		StatusCode: 0,
		StatusMsg:  "",
		VideoList:  videoList,
	})
}
