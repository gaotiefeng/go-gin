package formatter

import "go-gin/app/model"

func UserBase (user model.User) map[string]interface{}{
	return map[string]interface{}{
		"id" :	user.Id,
		"mobile" : user.Mobile,
		"password" : user.Password,
	}
}
