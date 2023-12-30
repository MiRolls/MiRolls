package mainProgram

import "github.com/gin-gonic/gin"

func MiddleWare(c *gin.Context) {
	c.Header("Author", "Liangmi and  Kovloen")
	c.Next()
}
