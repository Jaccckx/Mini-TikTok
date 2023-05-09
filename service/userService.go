package service

type UserService interface {
	LikeService
	FollowService

	GetUserInfoById(id uint64) (*User, error)

	GetUserInfoListById(id []uint64) ([]*User, error)

	GetUserInfoByName(name string) (*User, error)

	InsertUser(user *User) error
}
