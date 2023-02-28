package routes

import "github.com/gin-gonic/gin"

func DownloadAndGetDownloadSpeed(r *gin.Engine) {
	r.POST("/install/")
}
