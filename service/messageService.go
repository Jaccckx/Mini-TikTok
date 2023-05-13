package service

import "mini-tiktok/dao"

type MessageService interface {
	SendMessage(fromUserID int64, toUserID int64, content string) error
	GetMessageList(fromUserID int64, toUserID int64, preMsgTime int64) ([]*dao.Message, error)
}
