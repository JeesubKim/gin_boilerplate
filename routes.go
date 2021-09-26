package main

import (
	"github.com/gin-gonic/gin"
)

func initRoutes() *gin.Engine {
	r := gin.Default()
	initPageRouter(r)
	initUserRouter(r)
	return r
}

func initPageRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"What is": "gin.H?",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"login": "api",
		})
	})
	r.GET("/signup", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"signup": "api",
		})
	})
	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"code":    "PAGE_NOT_FOUND",
	// 		"message": "Page not found",
	// 	})
	// })
}

func initUserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/login", userLogin)
		user.GET("/logout", userLogout)
		user.POST("/register", userRegister)
	}
}
