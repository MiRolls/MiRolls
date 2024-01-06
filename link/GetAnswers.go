package link

import (
	"MiRolls/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type codeInfo struct {
	Code string `json:"code"`
}

func GetAnswers(c *gin.Context) {
	body, _ := c.GetRawData()
	codeStruct := new(codeInfo)
	_ = json.Unmarshal(body, &codeStruct) // Unmarshal the data to codeStruct from the frontend
	code := codeStruct.Code

	var answerList []interface{}
	err := database.Db.Select(&answerList, "SELECT `answer` FROM `answer` WHERE `code`=?", code)
	if err != nil {
		c.JSON(501, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"answers": answerList,
	})
}
