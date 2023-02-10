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
	isSuccess, errCode := config.InitConfig()
	if !isSuccess && errCode == 0 {
		//is install mode
		r := gin.Default()
		path, _ := filepath.Abs("install/routes")
		r.Static("/", path)
		//Load static files
	}
	//Install

	//gin.SetMode(gin.ReleaseMode)
	// set release
	r := gin.Default()

	//Load MiddleWare
	r.Use(MiddleWare)

	//Load statics
	path, _ := filepath.Abs(config.Configs.Server.Static)
	r.Static("/", path)

	//Register Router
	link.GetSite(r)
	link.QueryRoll(r)
	link.CreateRoll(r)
	link.NotFound(r)

	err := r.Run(":" + fmt.Sprintf("%d", config.Configs.Server.Port))
	if err != nil {
		log.Fatal("[FATAL ERROR]Cannot start server")
	} else {
		log.Println("[Success]Server running at http://" + config.Configs.Site.Link + ":" + fmt.Sprintf("%d", config.Configs.Server.Port) + "/")
	}
}
