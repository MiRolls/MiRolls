package server

import (
	"MiRolls/config"
	"MiRolls/install/routes"
	"MiRolls/link"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
)

func Boot() {
	isSuccess, errCode := config.InitConfig()
	if !isSuccess && errCode == 0 {
		log.Println("[Warning]MiRolls can't find config.yaml, It's Running the Install Mode. Server run at localhost:2333")
		//is install mode
		r := gin.Default()
		path, _ := filepath.Abs("./install/static/dist")
		r.Static("/", path)
		//Load static files
		routes.SetSite(r)
		routes.SetDatabase(r)
		routes.DownloadAndGetDownloadSpeed(r)
		r.NoRoute(func(context *gin.Context) {
			context.File("./install/static/dist/index.html")
		})
		//link.NotFound(r)
		//Load routes
		_ = r.Run(":2333")
		return
	}
	//Install

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
		//goland:noinspection HttpUrlsUsage
		log.Println("[Success]Server running at http://" + config.Configs.Site.Link + ":" + fmt.Sprintf("%d", config.Configs.Server.Port) + "/")
	}
}
