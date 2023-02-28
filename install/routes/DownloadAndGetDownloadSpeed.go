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
			err = DownloadFile("./theme", "")
			if err != nil {
				return
			}
		} else {
			c.JSON(500, gin.H{
				"message": "error",
				"error":   "Are you idiot? What are you doing?" + err.Error(),
			})
		}
	})
}

func DownloadFile(filepath string, url string) error {
	// create a file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err = out.Close()
		if err != nil {

		}
	}(out)

	// Get respone
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// write file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
