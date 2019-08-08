package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Base struct {
	c *gin.Context
}

func (a *Base) Api(c *gin.Context, code int, data interface{})  {
	c.JSON(code, data)
}

func (s *Base) Success(c *gin.Context, msg string, data ...interface{}){
	var d interface{}
	if data != nil && len(data) > 0 {
		d = data[0]
	}
	s.Api(c, http.StatusOK, gin.H{
		"success" : true,
		"msg" : msg,
		"data" : d,
	})
}

func (e *Base) Error(c *gin.Context, msg string, ErrorCode uint)  {
	e.Api(c,http.StatusOK,gin.H{
		"success" : false,
		"msg" : msg,
		"ErrorCode" : ErrorCode,
	})
}