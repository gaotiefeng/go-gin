package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	msg := fmt.Sprintf("welcome to mobile %d password %d", mobile, password)

	c.JSON(http.StatusOK,gin.H{
		"msg" : msg,
	})
}

func UserInfo(c *gin.Context) {

	id := c.Param("id")

	c.String(http.StatusOK,"Your message !!! Your id is %d", id)
}