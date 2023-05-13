package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mini-tiktok/service"
	"strconv"
)

func SendMessage(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("[SendMessage] error ", msg)
		c.JSON(200, MessageActionResponse{
			StatusCode: code,
			StatusMsg:  msg,
		})
	}

	fromUserID, err := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	if err != nil {
		errHan(400, "user_id empty")
		return
	}

	toUserID, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		errHan(400, "to_user_id empty")
		return
	}

	if fromUserID == toUserID {
		errHan(400, "from_user_id == to_user_id")
		return
	}

	content := c.Query("content")
	if content == "" {
		errHan(400, "content empty")
		return
	}

	msi := service.MessageServiceImpl{}
	err = msi.SendMessage(fromUserID, toUserID, content)
	if err != nil {
		errHan(500, err.Error())
		return
	}

	c.JSON(200, MessageActionResponse{
		StatusCode: 200,
		StatusMsg:  "success",
	})
}

func GetMessageChat(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("[SendMessage] error ", msg)
		c.JSON(200, MessageActionResponse{
			StatusCode: code,
			StatusMsg:  msg,
		})
	}

	fromUserID, err := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	if err != nil {
		errHan(400, "user_id empty")
		return
	}

	toUserID, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		errHan(400, "to_user_id empty")
		return
	}

	if fromUserID == toUserID {
		errHan(400, "from_user_id == to_user_id")
		return
	}

	preMsgTime, err := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	if err != nil {
		preMsgTime = 0
	}

	msi := service.MessageServiceImpl{}
	list, err := msi.GetMessageList(toUserID, fromUserID, preMsgTime)
	if err != nil {
		errHan(500, err.Error())
		return
	}

	c.JSON(200, MessageListResponse{
		MessageList: list,
		StatusCode:  0,
		StatusMsg:   "",
	})

}
