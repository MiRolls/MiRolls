package link

import (
	"MiRolls/database"
	"MiRolls/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
)

// CreateQuestionnaire create a questionnaire with `questionnaire` structure
func CreateQuestionnaire(c *gin.Context) {
	// Firstly,  Get request body
	reqBody, _ := c.GetRawData()

	// Secondly, Check the request
	Questionnaire := new(utils.Questionnaire)
	err := json.Unmarshal(reqBody, &Questionnaire)
	if err != nil { // check not success
		c.JSON(400, gin.H{
			"message": "error",
			"error":   "it isn't questionnaire",
		})
	}

	// Thirdly, Write database
	code := utils.Md5Hash(string(reqBody) + strconv.Itoa(rand.Int()))
	link := utils.Md5Hash(code)
	_, err = database.Db.Exec("INSERT INTO `rolls`(`id`,`roll`,`code`,`link`) VALUES(DEFAULT,?,?,?)", string(reqBody), code, utils.Md5Hash(code))
	if err != nil {
		c.JSON(503, gin.H{
			"message": "error",
			"error":   "cannot write in database",
		})
		return
	}

	// Fourthly, return
	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"code": code,
			"link": link,
		},
	})
}
