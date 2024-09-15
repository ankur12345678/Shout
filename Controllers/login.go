package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (base *BaseController) LoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "at login",
	})
}
