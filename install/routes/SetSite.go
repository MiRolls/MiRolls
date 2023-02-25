package routes

import (
	"MiRolls/config"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Site struct {
	Name      string `yaml:"name"`
	Link      string `yaml:"link"`
	Logo      string `yaml:"logo"`
	MainColor string `yaml:"mainColor"`
	Icp       string `yaml:"icp"`
	Lang      string `yaml:"lang"`
	NeedIcp   int    `yaml:"needIcp"`
}

func SetSite(r *gin.Engine) {
	r.POST("/install/set/site", func(c *gin.Context) {
		err := config.MakeConfig()
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
		}
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
			return
		}
		siteInfo := new(Site)
		err = json.Unmarshal(data, &siteInfo)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
			return
		}
		yamlConfig, err := yaml.Marshal(&siteInfo)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
		}
		err = config.ChangeConfig(1, string(yamlConfig))
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
