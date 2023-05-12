package main

import (
	"mini-tiktok/controller"
	"mini-tiktok/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// user api
	api := r.Group("/douyin")
	api.GET("/user/", jwt.Auth(), controller.GetUserInfo)
	api.POST("/user/register/", controller.Register)
	api.POST("/user/login/", controller.Login)
	api.GET("/feed", jwt.AuthNoLogin(), controller.Feed)
	//BUG:无法检测到token
	api.POST("/publish/action/", jwt.Auth(), controller.Publish)
}
