package server

import (
	"MiRolls/config"
	"MiRolls/database"
	"MiRolls/link"
	"MiRolls/packages"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"path/filepath"
)

func Boot() {
	isSuccess, errCode := config.InitConfig()
	if !isSuccess && errCode == 0 {
		//Install
		packages.RunSetupMode()
	}

	err, db := database.Open()
	if err != nil {
		log.Fatal("[FATAL]Database connection pool couldn't run.")
	}
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	r := gin.Default()
	//Load MiddleWare
	r.Use(MiddleWare)
	//Load statics
	path, _ := filepath.Abs(config.Configs.Server.Static)
	r.Static("/", path)

	//Register Router
	link.GetSite(r)
	link.QueryRoll(r, db)
	link.CreateRoll(r, db)
	link.NotFound(r)
	link.AnswerQuestionnaire(r, db)
	link.GetRoll(r, db)

	log.Println("[Success]Server running at " + config.Configs.Site.Link + ":" + fmt.Sprintf("%d", config.Configs.Server.Port) + "/")
	err = r.Run(":" + fmt.Sprintf("%d", config.Configs.Server.Port))
	if err != nil {
		log.Fatal("[FATAL ERROR]Cannot start server")
	}
}
