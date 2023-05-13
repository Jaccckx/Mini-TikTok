package service

import (
	"mini-tiktok/dao"
	"sort"
	"time"
)

type MessageServiceImpl struct{}

func (m *MessageServiceImpl) SendMessage(fromUserID int64, toUserID int64, content string) error {
	message := &dao.Message{
		Content:    content,
		CreateTime: time.Now().Unix(),
		FromUserID: fromUserID,
		ToUserID:   toUserID,
	}
	return dao.InsertMessage(message)
}

func (m *MessageServiceImpl) GetMessageList(fromUserID int64, toUserID int64, preMsgTime int64) ([]*dao.Message, error) {
	list1, err := dao.GetMessageList(fromUserID, toUserID, preMsgTime)
	if err != nil {
		return nil, err
	}

	list2, err := dao.GetMessageList(toUserID, fromUserID, preMsgTime)
	if err != nil {
		return nil, err
	}

	list := append(list1, list2...)
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreateTime < list[j].CreateTime
	})
	return list, nil
}
