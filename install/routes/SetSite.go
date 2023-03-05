package routes

import (
	"MiRolls/config"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func SetSite(r *gin.Engine) {
	r.POST("/install/set/site", func(c *gin.Context) {
		err := config.MakeConfig()
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
			return
		}
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
			return
		}
		siteInfo := new(config.Site)
		err = json.Unmarshal(data, &siteInfo)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
			return
		}
		//yamlConfig, err := yaml.Marshal(&siteInfo)
		//if err != nil {
		//	c.JSON(500, gin.H{
		//		"error":   err.Error(),
		//		"message": "error",
		//	})
		//}
		err = config.ChangeSite(siteInfo)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})
}
