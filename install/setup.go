package install

import (
	"MiRolls/config"
	"MiRolls/database"
	"MiRolls/mainProgram"
	"MiRolls/utils"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Println("[Warning]MiRolls can't find config.yaml, It's running the Install Mode. Server is going to run at localhost:2333")
	//read config.yaml and download file

	_, err := os.ReadDir("install")
	if err != nil {
		log.Println("[Warning]MiRolls can't find MiRolls-installer. Downloading MiRolls-installer now.")
		err = DownloadFile("./install/install.zip", config.CloudConfigs.InstallerDownload, "install")
		if err != nil {
			log.Fatal("Can't download files")
		}
		err = utils.Unzip("./install/install.zip", "./install")
		if err != nil {
			log.Fatal("Cant unzip.")
		}
	}
	//is installer
	r := gin.Default()
	path, _ := filepath.Abs("./install")
	r.Static("/", path)
	//Load static files
	SetSite(r)
	SetDatabase(r)
	Download(r, cancel)
	r.NoRoute(func(context *gin.Context) {
		context.File("./install/index.html")
	})
	//link.NotFound(r)
	//Load routes
	//_ = r.Run(":2333")
	//fmt.Println(111)
	//go func() { time.Sleep(time.Second * 5); fmt.Println(111) }()

	srv := &http.Server{
		Addr:    ":2333",
		Handler: r,
	}

	go func() {
		//srv.ListenAndServe
		err = srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("[FATAL] Can't run server." + err.Error())
		}
	}()
	log.Println("Install MiRolls at http://localhost:2333/")

	<-ctx.Done()
	log.Println("Shutting down server...")
	err = srv.Close()
	if err != nil {
		log.Fatal("[FATAL] Can't close server" + err.Error())
	}

	config.InitConfig()
	database.Open()
	mainProgram.Run()
	return
}
