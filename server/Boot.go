package server

import (
	"MiRollsBackend/config"
	"MiRollsBackend/link"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func Boot() {
	r := gin.Default()
	//处理中间件
	r.Use(MiddleWare)
	//加载html文件
	path, _ := filepath.Abs(config.FilePath)
	r.Static("/", path)
	//加载路由
	link.CreateRoll(r)
	link.NotFound(r)

	_ = r.Run(config.Post)
}
