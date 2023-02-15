package link

import (
	"MiRolls/config"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strconv"
)

func CreateRoll(r *gin.Engine) {
	r.POST("/roll/create", func(c *gin.Context) {
		mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Configs.Database.Username, config.Configs.Database.Password, config.Configs.Database.Protocol, config.Configs.Database.Host, config.Configs.Database.Port, config.Configs.Database.Database)
		sql, err := sqlx.Open("mysql", mysql)
		if err != nil {
			c.JSON(500, gin.H{
				"message":   "error",
				"error":     err.Error(),
				"errorType": "DataBase connect error",
			})
		}

		// get requestBody
		reqBody, _ := c.GetRawData()
		//reqBody := "1"
		code := md5Hash(string(reqBody) + strconv.Itoa(rand.Int()))
		link := "https://" + config.Configs.Site.Link + "/#/query?code=" + md5Hash(code)
		// directly into database
		//goland:noinspection SqlResolve
		_, err = sql.Exec("INSERT INTO `rolls`(`id`,`roll`,`code`,`link`) VALUES(DEFAULT,?,?,?)", string(reqBody), code, md5Hash(code))
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

func md5Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
