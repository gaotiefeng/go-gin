package service

import (
	"errors"
	"go-gin/app/model"
	"go-gin/database"
)

type UserService struct {

}

func (u *UserService) UserAdd (user model.User) (model.User, error){
	db := database.NewDb()
	db.Create(&user)

	if db.NewRecord(user) == true	{
		return user, errors.New("创建用户失败")
	}

	return user,nil
}

func (u *UserService) UserInfo (id interface{}) (user model.User,err error)  {
	database.NewDb().Where("id = ?", id).First(&user)

	if user.Id == 0 {
		err = errors.New("未查到用户")
	}

	return user, err
}

func (u *UserService) UserLogin (mobile interface{},password interface{}) (user model.User,err error) {
	database.NewDb().Where("mobile = ? and password = ?",mobile,password).Find(&user)

	if user.Mobile == "" {
		err = errors.New("未查到用户")
	}

	return user, err
}