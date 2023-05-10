package service

import (
	"mini-tiktok/dao"
	"time"
)

type ActionType uint

const (
	ActionConfirm = iota
	ActionCancel
)

type User struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

type Comment struct {
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id"`          // 评论id
	User       *User  `json:"user"`        // 评论用户信息
}

func ToUser(user *dao.User) *User {
	return &User{
		FavoriteCount:   0,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		TotalFavorited:  "1",
		WorkCount:       0,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		ID:              user.ID,
		Name:            user.Name,
		Signature:       user.Signature,
	}
}

func ToComment(comment *dao.Comment) *Comment {
	return &Comment{
		Content:    comment.Content,
		CreateDate: time.Unix(comment.CommitTime, 0).Format("2006-01-02 15:04:"),
		ID:         comment.ID,
		User:       nil,
	}
}
