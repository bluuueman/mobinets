package router

import (
	"test/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/upload", service.FileUpload)
	router.POST("/execcmd", service.ExecCommand)
	router.GET("/docker/image", service.GetDockerImages)
	return router
}
