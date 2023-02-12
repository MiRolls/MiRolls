package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Server   *Server     `yaml:"server"`
	Database *Database   `yaml:"database"`
	Site     *SiteConfig `yaml:"site"`
}

type Server struct {
	Port   int    `yaml:"port"`
	Static string `yaml:"static"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type SiteConfig struct {
	Name      string `yaml:"name"`
	Link      string `yaml:"link"`
	Logo      string `yaml:"logo"`
	MainColor string `yaml:"mainColor"`
	Icp       string `yaml:"icp"`
	Lang      string `yaml:"lang"`
	NeedIcp   int    `yaml:"needIcp"`
}

type Site struct {
	Name    string
	Link    string
	Logo    string
	Icp     string
	Lang    string
	NeedIcp int
}

func SetSite(r *gin.Engine) {
	r.POST("/install/set/site", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
		}
		var siteInfo *Site
		err = json.Unmarshal(data, &siteInfo)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
		}
		//config.MakeConfig()
	})
}
