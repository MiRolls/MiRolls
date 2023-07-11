package install

import (
	"MiRolls/config"
	"MiRolls/packages"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

type file struct {
	File string `json:"file"`
}

type GithubApi struct {
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

var fileName = "./theme/default.zip"

func Download(r *gin.Engine, closeServer context.CancelFunc) {
	r.POST("/install/download", func(c *gin.Context) {
		cfg := new(config.Server)
		cfg.Port = 2333
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
			req, err := http.Get("https://api.github.com/repos/MiRolls/MiRolls-default-theme/releases/latest")
			defer func(Body io.ReadCloser) { _ = Body.Close() }(req.Body)

			responseJson, err := io.ReadAll(req.Body)
			if err != nil {
				c.JSON(500, gin.H{"message": "error", "error": err.Error()})
				return
			}
			gitHubApiResponse := new(GithubApi)
			_ = json.Unmarshal(responseJson, &gitHubApiResponse)
			// Link to default
			if err != nil {
				c.JSON(500, gin.H{
					"message": "error",
					"error":   "Can't get mirolls-default-theme" + err.Error(),
				})
				return
			}
			assets := gitHubApiResponse.Assets[0]
			err = DownloadFile(fileName, assets.BrowserDownloadUrl, "theme")
			// Download file
			if err != nil {
				return
			}
			err = unZip()
			if err != nil {
				return
			} // Unzip
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
	err := packages.Unzip(fileName, "./theme/")
	if err != nil {
		return err
	}
	return nil
}
