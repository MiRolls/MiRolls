package server

import (
	"MiRollsBackend/config"
	"MiRollsBackend/link"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func Boot() {
	r := gin.Default()
	//处理中间件
	r.Use(MiddleWare)
	//加载html文件
	path, _ := filepath.Abs(config.Configs.Server.Static)
	r.Static("/", path)
	//加载路由
	link.CreateRoll(r)
	link.NotFound(r)

	_ = r.Run(":" + fmt.Sprintf("%d", config.Configs.Server.Port))
}
