package mainProgram

import (
	"MiRolls/link"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/site/get", link.GetSite)
	r.POST("/bigdata/get", link.QueryRoll)
	r.POST("/questionnaire/create", link.CreateQuestionnaire)
	r.NoRoute(link.NotFound)
	r.POST("/answer/create", link.AnswerQuestionnaire)
	r.POST("/questionnaire/get", link.GetRoll)
	r.POST("/answer/get", link.GetAnswers)
}
