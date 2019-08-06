package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/app/model"
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

	ra, err := user.AddUser()
	if err != nil {
		log.Fatal(err)
	}
	msg := fmt.Sprintf("welcome to mobile %d password %d Id %d", mobile, password, ra)

	c.JSON(http.StatusOK,gin.H{
		"msg" : msg,
	})
}

func UserInfo(c *gin.Context) {

	id := c.Param("id")

	c.String(http.StatusOK,"Your message !!! Your id is %d", id)
}