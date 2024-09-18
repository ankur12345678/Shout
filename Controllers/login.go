package controllers

import (
	"net/http"

	models "github.com/ankur12345678/shout/Models"
	"github.com/ankur12345678/shout/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (base *BaseController) LoginHandler(c *gin.Context) {
	//TODO: here we assumed that user will pass email and password. write validation for this
	var reqUser models.User
	c.BindJSON(&reqUser)
	//check if user present in the db
	userRepo := models.InitUserRepo(Ctrl.DB)
	user, err := userRepo.GetByEmail(reqUser.Email)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "no records found, please checkout signup!",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "error getting user details from DB",
		})
		return
	}
	//verify password
	match := utils.VerifyPassword(reqUser.Password, user.Password)
	if !match {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "invalid password!",
		})
		return
	}
	//generate a jwt with time 10 min
	accessToken, err := utils.GenerateJWT(Ctrl.Config.JWT_SECRET, reqUser.Email, Ctrl.Config.JWT_EXPIRY_TIME)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "error generating access token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"expiresIn":    Ctrl.Config.JWT_EXPIRY_TIME,
	})

}
