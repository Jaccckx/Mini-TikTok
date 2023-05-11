package dao

type Comment struct {
	ID         int64
	UserID     int64
	VideoID    int64
	Content    string
	CommitTime int64
}

func (Comment) TableName() string {
	return "comment" // 自定义表名
}

// InsertComment 插入数据, 返回 comment id
func InsertComment(comment *Comment) (int64, error) {
	result := Db.Create(comment)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return comment.ID, nil
}

// DeleteCommentByID 删除指定 id 的 comment
func DeleteCommentByID(id int64) error {
	result := Db.Delete(&Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetCommentListByVideoID 返回对应的 video list
func GetCommentListByVideoID(id int64) ([]*Comment, error) {
	var comments []*Comment
	result := Db.Where("video_id = ?", id).Take(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}
