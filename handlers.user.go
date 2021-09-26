package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userLogin(c *gin.Context) {
	var login user
	c.BindJSON(&login)
	c.JSON(http.StatusOK, gin.H{
		"id":       login.ID,
		"password": login.Password,
	})
}
func userLogout(c *gin.Context) {
	c.JSON(200, gin.H{
		"What is": "gin.H?",
	})
}
func userRegister(c *gin.Context) {
	c.JSON(200, gin.H{
		"What is": "gin.H?",
	})
}
