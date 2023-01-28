package server

import (
	"MiRolls/config"
	"MiRolls/link"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
)

func Boot() {
	gin.SetMode(gin.ReleaseMode)
	// set release
	r := gin.Default()
	//处理中间件
	r.Use(MiddleWare)
	//加载html文件
	path, _ := filepath.Abs(config.Configs.Server.Static)
	r.Static("/", path)
	//加载路由（是不是说成注册更好）
	link.CreateRoll(r)
	link.NotFound(r)
	link.GetSite(r)

	err := r.Run(":" + fmt.Sprintf("%d", config.Configs.Server.Port))
	if err != nil {
		log.Fatal("[FATAL ERROR]Cannot start server")
	} else {
		log.Println("[Success]Server running at http://localhost:" + fmt.Sprintf("%d", config.Configs.Server.Port) + "/")
	}
}
