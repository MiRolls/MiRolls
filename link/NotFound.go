package link

import (
	"MiRolls/config"
	"github.com/gin-gonic/gin"
)

func NotFound(r *gin.Engine) {
	r.NoRoute(func(context *gin.Context) {
		//context.HTML(404, "404.html", nil)
		context.File(config.Configs.Server.Static + "/index.html")
	})
}
