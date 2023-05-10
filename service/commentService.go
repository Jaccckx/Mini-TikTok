package service

type CommentService interface {
	GetCommentListByID(id uint64) ([]Comment, error)

	InsertComment(comment *Comment) error

	DeleteCommentByID(id uint64) error
}
