package routes

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func DownloadAndGetDownloadSpeed(r *gin.Engine) {
	r.POST("/install/download", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			return
		}
		if string(data) == "default" {
			DownloadFile("./theme", "")
		}
	})
}

func DownloadFile(filepath string, url string) error {
	// 创建一个空文件
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	// 获取文件响应
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 将响应内容写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
