package link

import (
	"MiRolls/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strconv"
)

func CreateRoll(r *gin.Engine, sql *sqlx.DB) {
	r.POST("/roll/create", func(c *gin.Context) {

		// get requestBody
		reqBody, _ := c.GetRawData()
		//reqBody := "1"
		code := utils.Md5Hash(string(reqBody) + strconv.Itoa(rand.Int()))
		//link := "https://" + config.Configs.Site.Link + "/#/query?code=" + md5Hash(code)
		link := utils.Md5Hash(code)
		// directly into database
		//goland:noinspection SqlResolve
		_, err := sql.Exec("INSERT INTO `rolls`(`id`,`roll`,`code`,`link`) VALUES(DEFAULT,?,?,?)", string(reqBody), code, utils.Md5Hash(code))
		if err != nil {
			c.JSON(500, gin.H{
				"message":   "error",
				"error":     err.Error(),
				"errorType": "DataBase insert error",
			})
			return
		}
		// All information has been placed in json. So we can directly into database
		// After placed in database, we'll return json of Frontend
		c.JSON(200, gin.H{
			"message":  "success",
			"rollCode": code,
			"rollLink": link,
		})
	})
	return
}
