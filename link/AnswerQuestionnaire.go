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

func AnswerQuestionnaire(r *gin.Engine, db *sqlx.DB) {
	r.POST("/answer", func(c *gin.Context) {
		body, _ := c.GetRawData()
		answerRel := new(answer)
		_ = json.Unmarshal(body, &answerRel)

		_, err := db.Exec("INSERT INTO `answer`(`id`,`anwser`,`link`) VALUES (DEFAULT,?,?)", answerRel.Answer, answerRel.Link)
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
