package controllers

import (
	"net/http"

	models "github.com/ankur12345678/shout/Models"
	"github.com/gin-gonic/gin"
)

func (base *BaseController) RootHandler(c *gin.Context) {
	//fetch top posts with highest engagment and show them
	var posts []models.Post
	err := Ctrl.DB.Order("likes DESC").Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "error fetching posts",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
