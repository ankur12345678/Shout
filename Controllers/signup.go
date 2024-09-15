package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (base *BaseController)SignUpHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "at sign up",
	})
}
