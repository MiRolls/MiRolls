package routes

import (
	"MiRolls/packages"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type file struct {
	File string `json:"file"`
}
type githubApi struct {
	Url       string `json:"url"`
	AssetsUrl string `json:"assets_url"`
	UploadUrl string `json:"upload_url"`
	HtmlUrl   string `json:"html_url"`
	Id        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeId          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []struct {
		Url      string      `json:"url"`
		Id       int         `json:"id"`
		NodeId   string      `json:"node_id"`
		Name     string      `json:"name"`
		Label    interface{} `json:"label"`
		Uploader struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadUrl string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballUrl string `json:"tarball_url"`
	ZipballUrl string `json:"zipball_url"`
	Body       string `json:"body"`
}

var downloadSpeed float64
var fileName = "./theme/default.zip"
var oldFileSize = 0.0
var fileSize = 0.0
var isDone = false
var unZipDone = false

func DownloadAndGetDownloadSpeed(r *gin.Engine) {
	r.POST("/install/download", func(c *gin.Context) {
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
		if string(data) == "default" {
			req, err := http.Get("https://api.github.com/repos/MiRolls/MiRolls-default-theme/releases/latest")
			defer func(Body io.ReadCloser) { _ = Body.Close() }(req.Body)

			responseJson, err := io.ReadAll(req.Body)
			if err != nil {
				c.JSON(500, gin.H{"message": "error", "error": err.Error()})
				return
			}
			gitHubApiResponse := new(githubApi)
			_ = json.Unmarshal(responseJson, &gitHubApiResponse)
			// Link to default
			if err != nil {
				c.JSON(500, gin.H{
					"message": "error",
					"error":   "Can't get mirolls-default-theme" + err.Error(),
				})
				return
			}
			err = DownloadFile(fileName, gitHubApiResponse.ZipballUrl)
			// Download file
			for !isDone {
				err := downloadSpeedControl()
				if err != nil {
					c.JSON(500, gin.H{
						"message": "error",
						"error":   err.Error(),
					})
					return
				}
				time.Sleep(1 * time.Second)
			}
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
	//Download api

	r.POST("/install/download/speed", func(c *gin.Context) {
		if fileSize == oldFileSize {
			isDone = true
			_ = unZip()
			c.JSON(200, gin.H{
				"done":    isDone,
				"message": "success",
				"speed":   "It's unzipping",
			})
		} else if unZipDone {
			c.JSON(200, gin.H{
				"done":    unZipDone,
				"message": "success",
				"speed":   "Done.",
			})
		} else {
			c.JSON(200, gin.H{
				"done":    isDone,
				"message": "success",
				"speed":   downloadSpeed,
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

	fileSize, err = strconv.ParseFloat(resp.Request.Header.Get("Content-Length"), 64)
	return nil
}

func downloadSpeedControl() error {
	stat, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	var newFileSize float64
	newFileSize = Round(float64(stat.Size() / 1048576)) //To mb/s
	downloadSpeed = newFileSize - oldFileSize
	oldFileSize = newFileSize
	//if fileSize
	return nil
}

func unZip() error {
	err := packages.Unzip(fileName, "./theme/")
	if err != nil {
		return err
	}
	return nil
}

func Round(number float64) float64 {
	return math.Round(number*100) / 100
}
