package controller

import (
	"log"
	"mini-tiktok/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// breif：从gin.Context请求里提取lastest_time参数和token参数，并返回视频列表
// @para c *gin.Context：HTTP请求上下文
//
// 函数内容：
// 1.提取lastest_time参数，并调用GetTime函数处理格式
// 2.提取token参数
// (1)如果token参数不为空，表明当前用户登录，需进行jwt授权检查。
// 3.根据lastest_time参数，返回视频列表。
// 4.如果视频列表为空，statusCode设为http.StatusNotFound，StatusMsg设为GetVideoFail
// BUG:点赞模块需要和用户模块联调
func Feed(c *gin.Context) {
	timeStr := c.Query("lastest_time")
	timeUnix := GetTime(timeStr)

	token := c.Request.Header.Get("Authorization")
	currUserID := int64(-1)
	if len(token) != 0 {
		//解析读取userId，并进行jwt校验
	}

	statusCode := int64(0)
	statusMsg := &getVideoSucc

	log.Println("GetVideosInfo")
	videoServiceTemp := service.VideoServiceImpl{}
	videoList, nextTime := videoServiceTemp.GetVideos(timeUnix, currUserID)
	if len(videoList) == 0 {
		statusCode = 1
		statusMsg = &emptyVideoErr
		log.Println("feed stream Get Empty VideoList")
	}
	c.JSON(200, FeedResponse{
		NextTime:   &nextTime,
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
		VideoList:  videoList,
	})

}

// breif：处理latest_time参数，使其由string格式变成time.Time格式
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
		log.Println("feed stream ParseTime failed")
		return
	}
	timeUnix = time.Unix(timeInt, 0)
	return
}
