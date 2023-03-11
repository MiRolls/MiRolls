package link

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type answer struct {
	Link   string `json:"link"`
	Answer string `json:"answer"`
}

func AnswerQuestionnaire(r *gin.Engine) {
	r.POST("/answer", func(c *gin.Context) {
		body, _ := c.GetRawData()
		answerRel := new(answer)
		_ = json.Unmarshal(body, answerRel)
	})
}
