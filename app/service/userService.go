package service

import (
	"errors"
	"go-gin/app/model"
	db "go-gin/database"
)

type UserService struct {

}

func (u *UserService) UserInfo (id interface{}) (user model.User,err error)  {

	db.SqlDB.Query("SELECT * FROM user WHERE id = ?", id)

	if user.Id == 0 {
		err = errors.New("未查到用户")
	}

	return user, err
}