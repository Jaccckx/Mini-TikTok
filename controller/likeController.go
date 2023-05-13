package controller

import (
	"mini-tiktok/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ActionLike(c *gin.Context) {

	currUserID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type := c.Query("action_type")

	StatusCode := int64(0)
	StatusMsg := ""

	LikeService := service.LikeServiceImpl{}
	IsLike, err := LikeService.IsFavoriteVideo(currUserID, videoID)
	if err == nil && ((IsLike == true && action_type == "1") || (IsLike == false && action_type == "2")) {
		StatusCode = 1
		StatusMsg = "操作失败"
	} else {
		LikeService.LikeVideo(currUserID, videoID, action_type)
	}
	c.JSON(200, FavoriteActionResponse{
		StatusCode: StatusCode,
		StatusMsg:  StatusMsg,
	})
}

func Favoritelist(c *gin.Context) {
	currUserID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	LikeService := service.VideoServiceImpl{}
	videoList := LikeService.GetLikeVideosList(currUserID)
	c.JSON(200, FavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "",
		VideoList:  videoList,
	})
}
