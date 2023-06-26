package middleware

import "github.com/gin-gonic/gin"

func GetUserAgent(c *gin.Context) string {
	return c.GetHeader("User-Agent")
}
