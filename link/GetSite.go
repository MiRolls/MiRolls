package link

import (
	. "MiRolls/config"
	"github.com/gin-gonic/gin"
)

func GetSite(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":      200,
		"message":   "success",
		"name":      Configs.Site.Name,
		"link":      Configs.Site.Link,
		"logo":      Configs.Site.Logo,
		"mainColor": Configs.Site.MainColor,
		"icp":       Configs.Site.Icp,
		"lang":      Configs.Site.Lang,
		"needIcp":   Configs.Site.NeedIcp,
	})
}
