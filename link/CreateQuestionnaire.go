package link

import (
	"MiRolls/database"
	"MiRolls/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
)

func CreateQuestionnaire(c *gin.Context) {
	// get requestBody
	reqBody, _ := c.GetRawData()
	//reqBody := "1"
	code := utils.Md5Hash(string(reqBody) + strconv.Itoa(rand.Int()))
	//link := "https://" + config.Configs.Site.Link + "/#/query?code=" + md5Hash(code)
	link := utils.Md5Hash(code)
	// directly into database
	//goland:noinspection SqlResolve
	_, err := database.Db.Exec("INSERT INTO `rolls`(`id`,`roll`,`code`,`link`) VALUES(DEFAULT,?,?,?)", string(reqBody), code, utils.Md5Hash(code))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	// All information has been placed in json. So we can directly into database
	// After placed in database, we'll return json of Frontend
	c.JSON(200, gin.H{
		"message": "success",
		"data": gin.H{
			"code": code,
			"link": link,
		},
	})
}
