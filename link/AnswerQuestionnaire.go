package link

import (
	"MiRolls/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type answer struct {
	Link   string `json:"link"`
	Answer string `json:"answer"`
}

func AnswerQuestionnaire(r *gin.Engine) {
	r.POST("/answer", func(c *gin.Context) {
		body, _ := c.GetRawData()
		answerRel := new(answer)
		_ = json.Unmarshal(body, &answerRel)
		mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Configs.Database.Username, config.Configs.Database.Password, config.Configs.Database.Protocol, config.Configs.Database.Host, config.Configs.Database.Port, config.Configs.Database.Database)
		db, err := sqlx.Open("mysql", mysql)
		if err != nil {
			c.JSON(501, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}
		_, err = db.Exec("INSERT INTO `answer`(`id`,`anwser`,`link`) VALUES (DEFAULT,?,?)", answerRel.Answer, answerRel.Link)
		// insert database
		if err != nil {
			c.JSON(500, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
		}) // return message
	})
}
