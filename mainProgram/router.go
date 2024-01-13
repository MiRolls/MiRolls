package mainProgram

import (
	"MiRolls/link"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.POST("/get/site", link.GetSite)
	r.POST("/query/roll", link.QueryRoll)
	r.POST("/questionnaire/create", link.CreateQuestionnaire)
	r.NoRoute(link.NotFound)
	r.POST("/answer", link.AnswerQuestionnaire)
	r.POST("/get/roll", link.GetRoll)
	r.POST("/get/answers", link.GetAnswers)
}
