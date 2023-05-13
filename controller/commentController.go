package controller

import (
	"mini-tiktok/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ActionComment(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type := c.Query("action_type")
	comment_text := c.Query("comment_text")
	comment_id, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

	commentService := service.CommentServiceImpl{}
	var commentID int64
	//添加评论
	if action_type == "1" {
		commentID, _ = commentService.CreateComment(userID, videoID, comment_text)
	} else if action_type == "2" {
		commentService.DeleteCommentByID(comment_id)
	}

	commentList, _ := commentService.GetCommentByCommentID(commentID, userID)
	c.JSON(200, CommentActionResponse{
		Comment:    commentList,
		StatusCode: 0,
		StatusMsg:  "",
	})
}

func CommentList(c *gin.Context) {
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	commentService := service.CommentServiceImpl{}
	commentList, _ := commentService.GetCommentListByVideoID(videoID, -1)
	c.JSON(200, CommentListResponse{
		CommentList: commentList,
		StatusCode:  0,
		StatusMsg:   "",
	})
}
