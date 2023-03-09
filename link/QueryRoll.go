package link

import (
	"MiRolls/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

func QueryRoll(r *gin.Engine) {
	r.POST("/query/roll", func(c *gin.Context) {
		mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Configs.Database.Username, config.Configs.Database.Password, config.Configs.Database.Protocol, config.Configs.Database.Host, config.Configs.Database.Port, config.Configs.Database.Database)
		sql, err := sqlx.Open("mysql", mysql)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error(), "error": err.Error(), "errorType": "database connect error"})
			log.Fatal("[FATAL ERROR]Cannot connect database")
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
		_ = sql.Close()
	})
}
