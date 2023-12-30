package link

import (
	"MiRolls/config"
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	//context.HTML(404, "404.html", nil)
	c.File(config.Configs.Server.Static + "/index.html")
}
