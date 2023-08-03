package link

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type answer struct {
	Link   string `json:"link"`
	Answer string `json:"answer"`
}

type rollFromDb struct {
	code string `db:"code"`
}

func AnswerQuestionnaire(r *gin.Engine, db *sqlx.DB) {
	r.POST("/answer", func(c *gin.Context) {
		body, _ := c.GetRawData()
		answerRel := new(answer)
		_ = json.Unmarshal(body, &answerRel)

		// Get "code" from "rolls" with "link"
		var ro rollFromDb // Get roll from database
		err := db.Get(&ro, "SELECT `code` from `answer` WHERE `link`=?", answerRel.Link)
		if err != nil {
			c.JSON(503, gin.H{
				"message": "error",
				"error":   err.Error(),
			})
			return
		}

		_, err = db.Exec("INSERT INTO `answer`(`id`,`answer`,`code`) VALUES (DEFAULT,?,?)", answerRel.Answer, ro.code)
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
