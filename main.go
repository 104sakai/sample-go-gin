package main

import (
	"net/http"

	"github.com/104sakai/sample-go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// middleware
	var ua string
	router.Use(func(c *gin.Context) {
		ua = middleware.GetUserAgent(c)
		c.Next()
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "pong",
			"User-Agent": ua,
		})
	})
	router.Run(":8080")
}
