package server

import "github.com/gin-gonic/gin"

func MiddleWare(c *gin.Context) {
	c.Header("author", "liangmi")
	c.Next()
}
