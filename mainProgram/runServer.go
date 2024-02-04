package mainProgram

import (
	"MiRolls/config"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
)

func Run() {
	r := gin.Default()
	//Load MiddleWare
	r.Use(GlobalMiddleWare())
	//Load statics
	path, _ := filepath.Abs(config.Configs.Server.Static)
	r.Static("/", path)

	//Register Router
	RegisterRouter(r)

	log.Println("[Success]Server is running, listen: " + config.Configs.Server.Bind)
	err := r.Run(config.Configs.Server.Bind)
	if err != nil {
		log.Fatal("[FATAL ERROR] Cannot start server")
	}
}
