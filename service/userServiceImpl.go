package service

import (
	"encoding/json"
	"fmt"
	"mini-tiktok/dao"
	"mini-tiktok/middleware/redis"
)

type UserServiceImpl struct {
	fsi FollowServiceImpl
	lsi LikeServiceImpl
	vsi VideoServiceImpl
}

func (u *UserServiceImpl) GetUserInfoById(userID int64, currUserID int64) (*User, error) {
	user, err := u.getUserInfo(userID)
	if err != nil {
		return nil, err
	}

	user.FavoriteCount, err = u.lsi.GetFavoriteCountByUserID(userID)
	if err != nil {
		return nil, err
	}

	user.FollowCount, err = u.fsi.GetFollowCount(userID)
	if err != nil {
		return nil, err
	}

	user.FollowerCount, err = u.fsi.GetFollowerCount(userID)
	if err != nil {
		return nil, err
	}

	if userID != currUserID {
		user.IsFollow, err = u.fsi.GetIsFollow(userID, currUserID)
		if err != nil {
			return nil, err
		}
	}

	user.WorkCount, err = u.vsi.GetVideoCount(userID)
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
	userDB := &dao.User{
		Name:            name,
		Password:        password,
		Avatar:          "https://siyuan-lieck.oss-cn-guangzhou.aliyuncs.com/v2-7bdf19ad23f3f04ae176680f0627fda5_xll.jpg",
		BackgroundImage: "https://siyuan-lieck.oss-cn-guangzhou.aliyuncs.com/v2-7bdf19ad23f3f04ae176680f0627fda5_xll.jpg",
		Signature:       "hhhhhhhhh",
	}

	id, err := dao.InsertUser(userDB)
	if err != nil {
		return 0, err
	}

	redisKey := fmt.Sprint("user_info_user_", id)
	user := ToUser(userDB)
	userStr, err := json.Marshal(user)
	if err != nil {
		return 0, err
	}

	err = redis.SetString(redisKey, string(userStr))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserServiceImpl) GetUserIdByName(name string) (int64, error) {
	return dao.GetUserIdByName(name)
}

func (u *UserServiceImpl) GetUserIdByNameAndPassword(name string, password string) (int64, error) {
	return dao.GetUserIdByPassword(name, password)
}

func (u *UserServiceImpl) getUserInfo(userID int64) (*User, error) {
	redisKey := fmt.Sprint("user_info_user_", userID)

	userInter, ok, err := redis.GetString(redisKey)
	if err != nil {
		return nil, err
	}

	// 缓存存在
	if ok {
		userInfo := User{}
		err := json.Unmarshal([]byte(userInter), &userInfo)
		if err != nil {
			return nil, err
		}

		return &userInfo, nil
	}

	// 不存在 写入缓存
	// 用户信息更改的频率低，因此可以直接获取缓存处理
	userDB, err := dao.GetUserInfoByID(userID)
	if err != nil {
		return nil, err
	}
	user := ToUser(userDB)

	userStr, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = redis.SetString(redisKey, string(userStr))
	if err != nil {
		return nil, err
	}

	return user, nil
}
