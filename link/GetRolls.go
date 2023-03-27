package link

import (
	"MiRolls/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type dataFromFront struct {
	Link string `json:"link,omitempty"`
	//Code string `json:"code"`
}

//type dber struct {
//	roll string `db:"roll"`
//}

func GetRoll(r *gin.Engine) {
	r.POST("/get/roll", func(c *gin.Context) {
		body, _ := c.GetRawData()
		dataFront := new(dataFromFront)
		_ = json.Unmarshal(body, &dataFront)

		mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Configs.Database.Username, config.Configs.Database.Password, config.Configs.Database.Protocol, config.Configs.Database.Host, config.Configs.Database.Port, config.Configs.Database.Database)
		db, err := sqlx.Open("mysql", mysql)
		if err != nil {
			c.JSON(501, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}
		defer func(db *sqlx.DB) {
			_ = db.Close()
		}(db)

		//roll := make(map[string]any)
		//roll := new(dber)
		roll := ""

		err = db.Get(&roll, "SELECT `roll` FROM `rolls` WHERE `link`=?", dataFront.Link)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
			"roll":    roll,
		}) // return message
	})
}
