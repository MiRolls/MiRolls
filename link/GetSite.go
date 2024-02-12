package link

import (
	"MiRolls/config"
	"github.com/gin-gonic/gin"
)

func GetSite(c *gin.Context) {
	c.JSON(200, config.Config.Key.Site)
}
