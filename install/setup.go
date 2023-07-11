package install

import (
	"MiRolls/mainProgram"
	"MiRolls/packages"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Println("[Warning]MiRolls can't find config.yaml, It's running the Install Mode. Server is going to run at localhost:2333")
	//read config.yaml and download file

	_, err := os.ReadDir("install")
	if err != nil {
		//err = os.Mkdir("install", 0777)
		//if err != nil {
		//	log.Fatal("[Error] Can't write files.")
		//}
		log.Println("[Warning]MiRolls can't find MiRolls-installer. Downloading MiRolls-installer now.")
		get, err := http.Get("https://api.github.com/repos/MiRolls/MiRolls-installer/releases/latest")
		if err != nil {
			log.Fatal("[Fatal] Cant get githubApi.")
		}
		log.Println("[Success]Got githubApi now.")
		responseJson, err := io.ReadAll(get.Body)
		if err != nil {
			log.Fatal("[Fatal] Cant to text.")
		}
		gitHubApiResponse := new(GithubApi)
		err = json.Unmarshal(responseJson, &gitHubApiResponse)
		if err != nil {
			log.Fatal("[FATAL] Cant unmarshal the github api response")
		}
		log.Println("[Success] Parse success. Download start.")
		err = DownloadFile("./install/install.zip", gitHubApiResponse.Assets[0].BrowserDownloadUrl, "install")
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

	go func(server *http.Server) {
		//srv.ListenAndServe
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("[FATAL] Can't run server")
		}
	}(srv)
	log.Println("Server is running at http://localhost:2333")
	err = srv.Shutdown(ctx)

	quit := make(chan os.Signal, 1)                      // 创建一个接收信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 让通道接收中断信号和终止信号

	<-quit

	fmt.Println("Shutting down server...")

	<-ctx.Done()
	err = srv.Close()
	if err != nil {
		log.Fatal("[FATAL] Can't close server" + err.Error())
	}

	mainProgram.Run()
	return
}
