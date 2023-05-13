package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mini-tiktok/service"
	"strconv"
)

// FollowAction 关注操作
func FollowAction(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("[FollowAction] error ", msg)
		c.JSON(200, RelationActionResponse{
			StatusCode: code,
			StatusMsg:  msg,
		})
	}

	userID, err := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	if err != nil {
		errHan(400, "user_id empty")
		return
	}

	toUserID, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		errHan(400, "to_user_id empty")
		return
	}

	if userID == toUserID {
		errHan(400, "to_user_id == user_id")
		return
	}

	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if actionType != 1 && actionType != 2 {
		errHan(400, "action_type empty")
		return
	}

	logrus.Debugf("[FollowAction] user_id:%v fools to_user_id: %v action_type: %v", userID, toUserID, actionType)

	fsi := service.FollowServiceImpl{}
	if actionType == 1 {
		err = fsi.FollowUser(userID, toUserID)
	} else {
		err = fsi.UnFollowUser(userID, toUserID)
	}

	if err != nil {
		errHan(500, err.Error())
		return
	}

	c.JSON(200, RelationActionResponse{
		StatusCode: 0,
		StatusMsg:  "",
	})

}

// FollowList 关注列表
func FollowList(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("follow list error ", msg)
		c.JSON(200, FollowListResponse{
			StatusCode: code,
			StatusMsg:  msg,
			UserList:   nil,
		})
	}

	currUserID, err := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	if err != nil {
		errHan(400, "curr user id empty")
		return
	}

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		errHan(400, "user id empty")
		return
	}

	logrus.Debugf("[FollowList] user_id:%v", userID)
	usi := service.UserServiceImpl{}
	fsi := service.FollowServiceImpl{}

	ids, err := fsi.GetFollowList(userID)
	if err != nil {
		errHan(500, err.Error())
		return
	}

	list, err := usi.GetUserInfoListById(ids, currUserID)
	if err != nil {
		errHan(500, err.Error())
		return
	}

	c.JSON(200, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "",
		UserList:   list,
	})
}

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("follower list error ", msg)
		c.JSON(200, FollowListResponse{
			StatusCode: code,
			StatusMsg:  msg,
			UserList:   nil,
		})
	}

	currUserID, err := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	if err != nil {
		errHan(400, "user_id error")
		return
	}

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		errHan(400, "user id empty")
		return
	}

	logrus.Debugf("[FollowerList] user_id:%v", userID)
	usi := service.UserServiceImpl{}
	fsi := service.FollowServiceImpl{}

	ids, err := fsi.GetFollowerList(userID)
	if err != nil {
		errHan(500, err.Error())
		return
	}

	list, err := usi.GetUserInfoListById(ids, currUserID)
	if err != nil {
		errHan(500, err.Error())
		return
	}

	c.JSON(200, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "",
		UserList:   list,
	})
}

// FriendList 好友列表
func FriendList(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("friend list error ", msg)
		c.JSON(200, FollowListResponse{
			StatusCode: code,
			StatusMsg:  msg,
			UserList:   nil,
		})
	}

	currUserID, err := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	if err != nil {
		errHan(400, "user_id error")
		return
	}

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		errHan(400, "user id empty")
		return
	}

	logrus.Debugf("[FriendList] user_id:%v", userID)
	usi := service.UserServiceImpl{}
	fsi := service.FollowServiceImpl{}

	ids, err := fsi.GetFriendList(userID)
	if err != nil {
		errHan(500, err.Error())
		return
	}

	list, err := usi.GetUserInfoListById(ids, currUserID)
	if err != nil {
		errHan(500, err.Error())
		return
	}

	c.JSON(200, FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "",
		UserList:   list,
	})
}
