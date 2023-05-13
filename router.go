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
	api.POST("/publish/action/", jwt.Auth(), controller.Publish)
	api.GET("/publish/list/", jwt.Auth(), controller.Publishlist)

	// follow api
	api.POST("/relation/action/", jwt.Auth(), controller.FollowAction)
	api.GET("/relation/follow/list/", jwt.Auth(), controller.FollowList)
	api.GET("/relation/follower/list/", jwt.Auth(), controller.FollowerList)
	api.GET("/relation/friend/list/", jwt.Auth(), controller.FriendList)
}
