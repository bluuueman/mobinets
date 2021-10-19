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
	router.DELETE("/k8s/service", service.DeleteService)
	router.POST("/k8s/service", service.StartService)
	router.DELETE("/k8s/deploy", service.DeleteDeploy)
	router.GET("/k8s/top", service.GetTop)
	return router
}
