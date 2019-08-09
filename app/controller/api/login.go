package api

import (
	"github.com/gin-gonic/gin"
	"go-gin/app/controller"
	"go-gin/app/form"
	"go-gin/app/formatter"
	"go-gin/app/model"
	"go-gin/app/service"
	errorCode "go-gin/config"
	"net/http"
)

type User struct {
	controller.Base
}

func Login(c *gin.Context)  {

	mobile := c.Request.FormValue("mobile")
	password := c.Request.FormValue("password")

	c.String(http.StatusOK, "request success mobile %d password %d", mobile, password)
}

func (u *User) Register(c *gin.Context)  {
	var (UserAddForm form.UserAddForm)
	err := c.BindJSON(&UserAddForm)
	if err != nil {
		u.Error(c,"json解析失败",500)
		return
	}

	err = u.Validator(c,UserAddForm)
	if err != nil {
		u.Error(c,"参数错误",600)
		return
	}

	user := model.User{Mobile:UserAddForm.Mobile, Password:UserAddForm.Password}
	_, err = new(service.UserService).UserAdd(user)
	if err != nil {
		u.Error(c,err.Error(),500)
		return
	}

	u.Success(c,"ok")
}

func (u *User) UserInfo(c *gin.Context) {

	id := c.Param("id")

	if 	id== "" || id == "0" {
		u.Error(c, "参数错误", errorCode.PARAMETER_ERROR)
		return
	}
	user, err := new(service.UserService).UserInfo(id)

	if err != nil {
		u.Error(c, "用户不存在", errorCode.USER_NOT_EXISTE)
		return
	}
	data := formatter.UserBase(user)

	u.Success(c, errorCode.SUCCESS, data)
}