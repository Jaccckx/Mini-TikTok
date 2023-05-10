package service

import "mini-tiktok/dao"

type UserServiceImpl struct {
	FollowServiceImpl
	LikeServiceImpl
	VideoServiceImpl
}

func (u *UserServiceImpl) GetUserInfoById(userID int64, currUserID int64) (*User, error) {
	userDb, err := dao.GetUserInfoByID(userID)
	if err != nil {
		return nil, err
	}
	user := ToUser(userDb)

	user.FavoriteCount, err = u.GetFavoriteCount(userID)
	if err != nil {
		return nil, err
	}

	user.FollowCount, err = u.GetFollowCount(userID)
	if err != nil {
		return nil, err
	}

	user.FollowerCount, err = u.GetFollowerCount(userID)
	if err != nil {
		return nil, err
	}

	if userID != currUserID {
		user.IsFollow, err = u.GetIsFollow(userID, currUserID)
		if err != nil {
			return nil, err
		}
	}

	user.WorkCount, err = u.GetVideoCount(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) GetUserInfoListById(ids []int64, currUserID int64) ([]*User, error) {
	var result []*User
	for _, id := range ids {
		user, err := u.GetUserInfoById(id, currUserID)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

func (u *UserServiceImpl) InsertUser(name string, password string) (int64, error) {
	return dao.InsertUser(&dao.User{
		Name:     name,
		Password: password,
	})
}

func (u *UserServiceImpl) GetUserIdByName(name string) (int64, error) {
	return dao.GetUserIdByName(name)
}

func (u *UserServiceImpl) GetUserIdByNameAndPassword(name string, password string) (int64, error) {
	return dao.GetUserIdByPassword(name, password)
}
