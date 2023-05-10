package service

type UserService interface {
	GetUserInfoById(userID int64, currUserID int64) (*User, error)

	GetUserInfoListById(id []int64, currUserID int64) ([]*User, error)

	InsertUser(name string, password string) (int64, error)

	GetUserIdByName(name string) (int64, error)

	GetUserIdByNameAndPassword(name string, password string) (int64, error)
}
