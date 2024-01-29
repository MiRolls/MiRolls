package link

import (
	. "MiRolls/config"
	"github.com/gin-gonic/gin"
)

// GetSite get "site" structure
func GetSite(c *gin.Context) {
	// Firstly, return site to frontend
	c.JSON(200, gin.H{
		"code":    200,
		"data":    Configs.Site,
		"message": "success",
	})
}
