package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPosts(c *gin.Context) {

}
func postCreate(c *gin.Context) {

}
func getPostDetail(c *gin.Context) {
	postID := c.Param("post_id")

	c.JSON(http.StatusOK, gin.H{
		"requestedID": postID,
	})
}
