package controllers

import (
	"net/http"

	models "github.com/ankur12345678/shout/Models"
	"github.com/gin-gonic/gin"
)

func (base *BaseController) UpdateBlogHandler(c *gin.Context) {
	var userRepo = models.InitUserRepo(Ctrl.DB)
	var postRepo = models.InitPostRepo(Ctrl.DB)
	var post models.Post
	c.BindJSON(&post)
	val, _ := c.Get("email")
	email := val.(string)
	user, err := userRepo.GetByEmail(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "error getting user",
		})
		return
	}
	postFromDB, err := postRepo.GetById(post.PostUUID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "trying to update invalid post",
		})
		return
	}
	if user.ID != uint(postFromDB.UserID) {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "unauthorized user",
		})
		return
	}
	// c.BindJSON(&post)
	err = postRepo.Update(&post, post.PostUUID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "error updating post",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
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
		c.JSON(http.StatusOK, gin.H{
			"error_message": "error fetching email details",
		})
		return
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
	reqestedId := c.Param("id")
	var postRepo = models.InitPostRepo(Ctrl.DB)
	post, err := postRepo.GetById(reqestedId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_message": "error fetching post details",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (base *BaseController) ShowAllBlogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "at blogs",
	})
}
