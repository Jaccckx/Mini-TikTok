package dao

type Message struct {
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	ID         int64  `json:"id"`           // 消息id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
}

func (Message) TableName() string {
	return "message"
}

func InsertMessage(message *Message) error {
	result := Db.Create(message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetMessageList(fromUserID int64, toUserID int64, preMsgTime int64) ([]*Message, error) {
	var messages []*Message
	result :=
		Db.Where("from_user_id =? AND to_user_id =? AND create_time > ?", fromUserID, toUserID, preMsgTime).
			Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}
