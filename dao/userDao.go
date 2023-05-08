package dao

type User struct {
	ID              int
	Name            string
	Password        string
	Avatar          string
	BackgroundImage string
	Signature       string
}

func InsertUser(user *User) {
	result := Db.Create(user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func GetUserInfoByID(id int) (*User, error) {
	var user User
	result := Db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserInfoByName(name string) (*User, error) {
	var user User
	result := Db.Where("name =?", name).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
