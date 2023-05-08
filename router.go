package main

import (
	"github.com/gin-gonic/gin"
	"mini-tiktok/controller"
	"mini-tiktok/middleware/jwt"
)

func initRouter(r *gin.Engine) {
	// user api
	api := r.Group("/douyin")
	api.GET("/user/", jwt.Auth(), controller.GetUserInfo)
	api.POST("/user/register/", controller.Register)
	api.POST("/user/login/", controller.Login)

}
