package install

import (
	"MiRolls/mainProgram"
	"MiRolls/packages"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
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
		log.Println("[Success]Got githubApi now.")
		responseJson, err := io.ReadAll(get.Body)
		if err != nil {
			log.Fatal("[Error] Cant to text.")
		}
		gitHubApiResponse := new(GithubApi)
		_ = json.Unmarshal(responseJson, &gitHubApiResponse)
		log.Println("[Success]Parse success. Download start.")
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
	cmd := exec.Command("r.Run", ":2333")
	go func(ctx context.Context, cmd *exec.Cmd) {
		//调用cmd.Start来启动进程
		if err := cmd.Start(); err != nil {
			log.Fatalf("start: %s\n", err)
		}
		//监听ctx.Done信号
		<-ctx.Done()
		//收到取消信号，结束进程
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("Process Kill:", err)
		}
	}(ctx, cmd)
	mainProgram.Run()
	return
}
