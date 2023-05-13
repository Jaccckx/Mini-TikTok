package controller

import "mini-tiktok/service"

/// 用户模块

type RegisterResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

type LoginResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

type UserInfoResponse struct {
	StatusCode int64         `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg"`  // 返回状态描述
	User       *service.User `json:"user"`        // 用户信息
}

/// 点赞模块

type FavoriteActionResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FavoriteListResponse struct {
	StatusCode string          `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string         `json:"status_msg"`  // 返回状态描述
	VideoList  []service.Video `json:"video_list"`  // 用户点赞视频列表
}

/// 评论模块

type CommentActionResponse struct {
	Comment    *service.Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int64            `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string          `json:"status_msg"`  // 返回状态描述
}

type CommentListResponse struct {
	CommentList []service.Comment `json:"comment_list"` // 评论列表
	StatusCode  int64             `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string           `json:"status_msg"`   // 返回状态描述
}

/// 关注模块

type RelationActionResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FollowListResponse struct {
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string          `json:"status_msg"`  // 返回状态描述
	UserList   []*service.User `json:"user_list"`   // 用户列表
}

/// 视频模块

var emptyVideoErr string = "GetEmptyVideo!"
var getVideoSucc string = "GetVideoSucc!"

type PublishActionResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type PublishListResponse struct {
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string          `json:"status_msg"`  // 返回状态描述
	VideoList  []service.Video `json:"video_list"`  // 用户发布的视频列表
}

type FeedResponse struct {
	NextTime   *int64          `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string         `json:"status_msg"`  // 返回状态描述
	VideoList  []service.Video `json:"video_list"`  // 视频列表
}
