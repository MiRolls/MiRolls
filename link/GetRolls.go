package link

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type dataFromFront struct {
	Link string `json:"link,omitempty"`
	//Code string `json:"code"`
}

func GetRoll(r *gin.Engine, db *sqlx.DB) {
	r.POST("/get/roll", func(c *gin.Context) {
		body, _ := c.GetRawData()
		dataFront := new(dataFromFront)
		_ = json.Unmarshal(body, &dataFront)
		roll := ""

		err := db.Get(&roll, "SELECT `roll` FROM `rolls` WHERE `link`=?", dataFront.Link)
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
