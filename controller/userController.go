package controller

import (
	"mini-tiktok/middleware/jwt"
	"mini-tiktok/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Register(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("user register error ", msg)
		c.JSON(200, RegisterResponse{
			StatusCode: code,
			StatusMsg:  msg,
			Token:      "",
			UserID:     0,
		})
	}

	name := c.Query("username")
	password := c.Query("password")

	if name == "" || password == "" {
		errHan(100, "username or password is empty")
		return
	}
	logrus.Debugf("user register name:%v pass:%v", name, password)

	usi := service.UserServiceImpl{}

	uid, err := usi.GetUserIdByName(name)
	if err != nil {
		errHan(100, err.Error())
		return
	}

	if uid != 0 {
		errHan(101, "用户名已存在")
		return
	}

	uid, err = usi.InsertUser(name, password)
	if err != nil {
		errHan(100, err.Error())
		return
	}

	token, err := jwt.NewToken(uid)

	if err != nil {
		errHan(100, err.Error())
		return
	}

	c.JSON(200, RegisterResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Token:      token,
		UserID:     uid,
	})
}

func Login(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("user login error ", msg)
		c.JSON(200, LoginResponse{
			StatusCode: code,
			StatusMsg:  msg,
			Token:      "",
			UserID:     0,
		})
	}

	name := c.Query("username")
	password := c.Query("password")

	usi := service.UserServiceImpl{}

	uid, err := usi.GetUserIdByNameAndPassword(name, password)
	if err != nil {
		errHan(100, err.Error())
		return
	}

	if uid == 0 {
		errHan(101, "用户名或密码错误")
		return
	}

	token, err := jwt.NewToken(uid)
	if err != nil {
		errHan(100, err.Error())
		return
	}

	c.JSON(200, LoginResponse{
		StatusCode: 0,
		StatusMsg:  "",
		Token:      token,
		UserID:     uid,
	})
}

func GetUserInfo(c *gin.Context) {
	errHan := func(code int64, msg string) {
		logrus.Error("user get user info error ", msg)
		c.JSON(200, UserInfoResponse{
			StatusCode: code,
			StatusMsg:  msg,
			User:       nil,
		})
	}

	uid, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		errHan(100, err.Error())
		return
	}

	usi := service.UserServiceImpl{}
	user, err := usi.GetUserInfoById(uid, uid)
	if err != nil {
		errHan(100, err.Error())
		return
	}

	c.JSON(200, UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "",
		User:       user,
	})
}
