package service

type CommentService interface {
	UserService

	GetCommentListByID(id uint64) ([]Comment, error)

	InsertComment(comment *Comment) error

	DeleteCommentByID(id uint64) error
}
