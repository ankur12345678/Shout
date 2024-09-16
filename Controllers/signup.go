package controllers

import (
	"net/http"

	models "github.com/ankur12345678/shout/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (base *BaseController) SignUpHandler(c *gin.Context) {
	var reqUser models.User
	c.ShouldBindJSON(&reqUser)
	userRepo := models.InitUserRepo(Ctrl.DB)
	//check if username exists previously
	_, err := userRepo.GetByUserName(reqUser.UserName)
	if err != gorm.ErrRecordNotFound  {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "username exists, please try to login",
		})
		return
	}
	//check if email exists previously
	_, err = userRepo.GetByEmail(reqUser.Email)
	if err != gorm.ErrRecordNotFound  {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "email exists, please try to login",
		})
		return
	}
	reqUser.UserUUID = UUIDGen("USER")
	err = userRepo.Create(&reqUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "error while creating user, please try again.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Created! Please visit login endpoint.",
	})
}
