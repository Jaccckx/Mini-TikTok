package service

type UserService interface {
	GetTableList() ([]string, error)
}
