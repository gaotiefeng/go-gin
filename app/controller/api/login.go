package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/app/model"
	"go-gin/app/service"
	"log"
	"net/http"
)

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

func UserInfo(c *gin.Context) {

	id := c.Param("id")

	fmt.Println("id = d%", id)

	if 	id== "" || id == "0" {
		c.String(http.StatusOK,"Your %d id empty", id)
	}
	user, err := new(service.UserService).UserInfo(id)

	if err != nil {
		c.String(http.StatusOK, "Your error %d", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg" : user,
	})
}