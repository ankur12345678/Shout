package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (base *BaseController) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "You are at home",
	})
}
