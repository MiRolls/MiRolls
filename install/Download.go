package install

import (
	"MiRolls/config"
	"MiRolls/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type file struct {
	File string `json:"file"`
}

var fileName = "./theme/default.zip"

func Download(r *gin.Engine, closeServer context.CancelFunc) {
	r.POST("/install/download", func(c *gin.Context) {
		var hasError *bool = new(bool)
		*hasError = true
		defer func(hasError *bool) {
			if *hasError == true {
				_ = config.Destroy()
			}
		}(hasError)

		cfg := new(config.Server)
		cfg.Bind = ":2333"
		cfg.Static = "theme"
		err := config.ChangeServer(cfg)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   err.Error(),
				"message": "error",
			})
			return
		}
		data, err := c.GetRawData()
		file := new(file)
		if json.Unmarshal(data, &file) != nil {
			c.JSON(500, gin.H{"message": "error", "error": err.Error()})
			return
		}
		if err != nil {
			return
		}
		// theme
		if file.File == "default" {
			err := DownloadFile(fileName, config.CloudConfigs.DefaultThemeDownload, "theme")
			// Download file
			if err != nil {
				return
			}
			err = unZip()
			if err != nil {
				return
			} // Unzip
			*hasError = false
			c.JSON(200, gin.H{
				"message": "success",
			})
		} else {
			c.JSON(500, gin.H{
				"message": "error",
				"error":   "Are you idiot? What are you doing?",
			})
		}
		closeServer()
	})

	//Download api
}

func DownloadFile(filepath string, url string, filepathName string) error {
	err := os.Mkdir(filepathName, 0755)
	if err != nil {
		return err
	}
	// create a file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	// Get response
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	// write file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func unZip() error {
	err := utils.Unzip(fileName, "./theme/")
	if err != nil {
		return err
	}
	return nil
}
