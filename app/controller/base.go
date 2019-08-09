package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smokezl/govalidators"
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

func (v *Base) Validator(c *gin.Context, data interface{})(err error) {
	validator := govalidators.New()
	err = validator.LazyValidate(data)
	if err != nil {
		fmt.Printf("err: %d",err)
		return	err
	}
	return nil
}