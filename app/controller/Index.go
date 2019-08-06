package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {

	c.String(http.StatusOK, "go to Index/Index")

}

func Welcome(c *gin.Context)  {

	c.String(http.StatusOK, "go to welcome")
}