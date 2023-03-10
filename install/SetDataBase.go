package install

import (
	"MiRolls/config"
	"MiRolls/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func SetDatabase(r *gin.Engine) {
	r.POST("/install/set/database", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		dbInfo := new(config.Database)
		err = json.Unmarshal(data, &dbInfo)
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		err = database.Init(dbInfo.Username, dbInfo.Password, dbInfo.Host, dbInfo.Database)
		if err != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		if config.ChangeDatabase(dbInfo) != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
}
