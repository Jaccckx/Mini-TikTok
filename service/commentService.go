package service

type CommentService interface {
	GetCommentListByVideoID(videoID int64, currUserID int64) ([]*Comment, error)

	CreateComment(userID int64, videoID int64, content string) (int64, error)

	DeleteCommentByID(id int64) error
}
