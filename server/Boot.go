package server

import (
	"MiRolls/config"
	"MiRolls/install"
	"MiRolls/link"
	"MiRolls/packages"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Boot() {
	isSuccess, errCode := config.InitConfig()

	if !isSuccess && errCode == 0 {
		log.Println("[Warning]MiRolls can't find config.yaml, It's running the Install Mode. Server is going to run at localhost:2333")
		//read config.yaml and download file

		_, err := os.ReadDir("install")
		if err != nil {
			err = os.Mkdir("install", 0777)
			if err != nil {
				log.Fatal("[Error] Can't write files.")
			}
			log.Println("[Warning]MiRolls can't find MiRolls-installer. Downloading MiRolls-installer now.")
			get, err := http.Get("https://api.github.com/repos/MiRolls/MiRolls-installer/releases/latest")
			log.Println("[Success]Got githubApi now.")
			responseJson, err := io.ReadAll(get.Body)
			if err != nil {
				log.Fatal("[Error] Cant to text.")
			}
			gitHubApiResponse := new(install.GithubApi)
			_ = json.Unmarshal(responseJson, &gitHubApiResponse)
			log.Println("[Success]Parse success. Download start.")
			err = install.DownloadFile("./install/install.zip", gitHubApiResponse.Assets[0].BrowserDownloadUrl, "install")
			if err != nil {
				log.Fatal("Can't download files")
			}
			err = packages.Unzip("./install/install.zip", "./install")
			if err != nil {
				log.Fatal("Cant unzip.")
			}
		}
		//is installer
		r := gin.Default()
		path, _ := filepath.Abs("./install")
		r.Static("/", path)
		//Load static files
		install.SetSite(r)
		install.SetDatabase(r)
		install.Download(r)
		r.NoRoute(func(context *gin.Context) {
			context.File("./install/index.html")
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
	link.AnswerQuestionnaire(r)

	err := r.Run(":" + fmt.Sprintf("%d", config.Configs.Server.Port))
	if err != nil {
		log.Fatal("[FATAL ERROR]Cannot start server")
	} else {
		//goland:noinspection HttpUrlsUsage
		log.Println("[Success]Server running at http://" + config.Configs.Site.Link + ":" + fmt.Sprintf("%d", config.Configs.Server.Port) + "/")
	}
}
