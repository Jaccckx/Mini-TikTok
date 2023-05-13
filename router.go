package main

import (
	"mini-tiktok/controller"
	"mini-tiktok/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	api := r.Group("/douyin")

	// user api
	api.GET("/user/", jwt.Auth(), controller.GetUserInfo)
	api.POST("/user/register/", controller.Register)
	api.POST("/user/login/", controller.Login)

	// video api
	api.GET("/feed", jwt.AuthNoLogin(), controller.Feed)
	api.POST("/publish/action/", jwt.Auth(), controller.Publish)
	api.GET("/publish/list/", jwt.AuthNoLogin(), controller.Publishlist)
	api.POST("/favorite/action/", jwt.Auth(), controller.ActionLike)
	api.GET("/favorite/list/", jwt.AuthNoLogin(), controller.Favoritelist)
	api.POST("/comment/action/", jwt.Auth(), controller.ActionComment)
	api.GET("/comment/list/", jwt.AuthNoLogin(), controller.CommentList)

	// follow api
	api.POST("/relation/action/", jwt.Auth(), controller.FollowAction)
	api.GET("/relation/follow/list/", jwt.Auth(), controller.FollowList)
	api.GET("/relation/follower/list/", jwt.Auth(), controller.FollowerList)
	api.GET("/relation/friend/list/", jwt.Auth(), controller.FriendList)

	// im api
	api.POST("/message/action/", jwt.Auth(), controller.SendMessage)
	api.GET("/message/chat/", jwt.Auth(), controller.GetMessageChat)
}
