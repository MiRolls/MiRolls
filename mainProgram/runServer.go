package mainProgram

import (
	"MiRolls/config"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	r := gin.Default()
	//Load MiddleWare
	r.Use(GlobalMiddleWare())
	//Load statics
	//path, _ := filepath.Abs(config.Config.Key..Static)
	//r.Static("/", path)

	//Register Router
	RegisterRouter(r)

	log.Println("[Success]Server is running, listen: " + config.Config.Key.Web.Address)
	err := r.Run(config.Config.Key.Web.Address)
	if err != nil {
		log.Fatal("[FATAL ERROR] Cannot start server")
	}
}
