package link

import (
	. "MiRolls/config"
	"github.com/gin-gonic/gin"
)

func GetSite(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"data":    Configs.Site,
		"message": "success",
	})
}
