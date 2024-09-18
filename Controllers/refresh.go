package controllers

import (
	"github.com/ankur12345678/shout/utils"
	"github.com/gin-gonic/gin"
)

func (base *BaseController) HandleRefresh(c *gin.Context) {
	//TODO: Blacklist the comming token and generate a new one
	email, exists := c.Get("email")
	if !exists {
		c.AbortWithStatusJSON(200, gin.H{
			"error_message": "error setting user in conetxt",
		})
		return
	}
	emailStr, ok := email.(string)
	if !ok {
		c.AbortWithStatusJSON(200, gin.H{
			"error_message": "error in type converison",
		})
		return
	}
	accessToken, err := utils.GenerateJWT(Ctrl.Config.JWT_SECRET, emailStr, Ctrl.Config.JWT_EXPIRY_TIME)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"error_message": "error in generating access token. Try again",
		})
		return
	}
	c.JSON(200, gin.H{
		"access_token": accessToken,
		"expiresIn":    600,
	})

}
