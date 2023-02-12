package routes

import (
	"github.com/gin-gonic/gin"
)

type Site struct {
	Name    string
	Link    string
	Logo    string
	Icp     string
	Lang    string
	NeedIcp bool
}

func SetSite(r *gin.Engine) {
	r.POST("/install/set/site", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(500, "Can't read reqBody")
		}

		//config.MakeConfig()
	})
}
