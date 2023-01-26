package link

import (
	. "MiRollsBackend/config"
	"github.com/gin-gonic/gin"
)

func GetSite(r *gin.Engine) {
	r.POST("/get/site", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":   "success",
			"name":      Configs.Site.Name,
			"link":      Configs.Site.Link,
			"logo":      Configs.Site.Logo,
			"mainColor": Configs.Site.MainColor,
			"icp":       Configs.Site.Icp,
		})
	})
}
