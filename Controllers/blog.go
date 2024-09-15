package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (base *BaseController) UpdateBlogHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "at blogs",
	})
}

func (base *BaseController) InsertBlogHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "at blogs",
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
