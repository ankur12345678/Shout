package controllers

import (
	"net/http"

	models "github.com/ankur12345678/shout/Models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (base *BaseController) UpdateBlogHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "at blogs",
	})
}

func (base *BaseController) InsertBlogHandler(c *gin.Context) {
	var userRepo = models.InitUserRepo(Ctrl.DB)
	var postRepo = models.InitPostRepo(Ctrl.DB)
	var post models.Post
	c.BindJSON(&post)
	post.Reputation = 0
	post.Share = 0

	//fetch the creater by email set in context
	val, _ := c.Get("email")
	userEmail := val.(string)
	//search in db with this email
	User, err := userRepo.GetByEmail(userEmail)
	if err != nil {
		log.Error("error fetching email details from user repo:")
	}
	post.UserID = int(User.ID)
	post.PostUUID = UUIDGen("POST")
	postRepo.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "post created",
		"post":    post,
	})
}

func (base *BaseController) ShowBlogById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "at blogs",
	})
}

func (base *BaseController) ShowAllBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "at blogs",
	})
}
