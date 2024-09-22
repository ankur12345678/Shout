package controllers

import (
	"net/http"
	"strconv"

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
	postFromDB, err := postRepo.GetById(post.ID)
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
	err = postRepo.Update(&post, post.ID)
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
	post.Likes = 0
	post.Replies = 0

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
	id, _ := strconv.ParseUint(reqestedId, 10, 64)
	var postRepo = models.InitPostRepo(Ctrl.DB)
	post, err := postRepo.GetById(uint(id))
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

func (base *BaseController) ShowMyBlogs(c *gin.Context) {
	//can use the email in the context to query the posts table
	//fetch the id of the email
	val, _ := c.Get("email")
	email := val.(string)
	userRepo := models.InitUserRepo(Ctrl.DB)
	user, err := userRepo.GetByEmail(email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error_message": "error fetching user detail",
		})
		return
	}
	var posts []models.Post

	Ctrl.DB.Model(&models.Post{}).Where(&models.Post{UserID: int(user.ID)}).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": "all of your posts",
		"posts":   posts,
	})
}

func (base *BaseController) DeleteBlogById(c *gin.Context) {
	//takes postID in body
	var requestedPost models.Post
	c.BindJSON(&requestedPost)
	

	postRepo := models.InitPostRepo(Ctrl.DB)
	post, err := postRepo.GetById(requestedPost.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error_message": "error fetching post ",
		})
		return
	}
	//fetch the user detail who is making the request
	val, _ := c.Get("email")
	userEmail := val.(string)
	userRepo := models.InitUserRepo(Ctrl.DB)
	user, err := userRepo.GetByEmail(userEmail)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error_message": "error fetching user",
		})
		return
	}
	if post.UserID != int(user.ID) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error_message": "unauthenticated user for the operation",
		})
		return
	}
	err = postRepo.Delete(post.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error_message": "error deleting the post",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post_uuid": post.PostUUID,
		"message":   "post deleted!",
	})
}

