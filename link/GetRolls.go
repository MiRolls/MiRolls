package link

import (
	"MiRolls/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type dataFromFront struct {
	Link string `json:"link,omitempty"`
	//Code string `json:"code"`
}

func GetRoll(c *gin.Context) {
	body, _ := c.GetRawData()
	dataFront := new(dataFromFront)
	_ = json.Unmarshal(body, &dataFront)
	roll := ""

	err := database.Db.Get(&roll, "SELECT `roll` FROM `rolls` WHERE `link`=?", dataFront.Link)
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
}
