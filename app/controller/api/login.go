package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/app/controller"
	"go-gin/app/model"
	"go-gin/app/service"
	errorCode "go-gin/config"
	"log"
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

func Register(c *gin.Context)  {

	mobile := c.Request.FormValue("mobile")
	password := c.Request.FormValue("password")

	user := model.User{Mobile:mobile, Password:password}

	_, err := new(service.UserService).UserAdd(user)

	if err != nil {
		log.Fatal(err)
	}
	msg := fmt.Sprintf("welcome to mobile %d password %d ", mobile, password)

	c.JSON(http.StatusOK,gin.H{
		"msg" : msg,
	})
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

	u.Success(c, errorCode.SUCCESS, user)
}