package service

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"test/utility"

	"github.com/gin-gonic/gin"
)

//文件上传
func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if utility.IsErr(err, "Upload File Failed!") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "upload file failed",
		})
		return
	}
	log.Println(file.Filename)
	dir := "tmp/"
	dst := fmt.Sprintf(dir + file.Filename)
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message":  "file upload succeed",
		"filepath": dst,
	})
}

//执行系统命令
func ExecCommand(c *gin.Context) {
	type msg struct {
		Message string `json:"message"`
	}
	jsondata := msg{}
	bindErr := c.BindJSON(&jsondata)
	if utility.IsErr(bindErr, "BindJSON Failed!") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server JSON bind failed",
		})
		return
	}
	target := "tmp/writefile.txt"
	cmd := exec.Command("/bin/bash", "-c", "echo"+" "+jsondata.Message+" > "+target)
	cmdRunErr := cmd.Run()
	if utility.IsErr(cmdRunErr, "Command Exec Failed!") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Command exec failed",
		})
		return
	}
	log.Println(jsondata.Message)
	c.JSON(http.StatusOK, gin.H{
		"message": "Command exec succeed",
		"target":  target,
	})
}

func StartService(c *gin.Context) {
	type msg struct {
		File string `json:"file"`
	}
	jsondata := msg{}
	bindErr := c.BindJSON(&jsondata)
	if utility.IsErr(bindErr, "BindJSON Failed!") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server JSON bind failed",
		})
		return
	}
	file := jsondata.File
	var output bytes.Buffer
	if !utility.KubeApplyYaml(file, &output) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Start service failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Service start succeed",
		"data":    output.String(),
	})
	return
}

func DeleteService(c *gin.Context) {
	type msg struct {
		Service   string `json:"service"`
		Namespace string `json:"namespace"`
	}
	jsondata := msg{}
	bindErr := c.BindJSON(&jsondata)
	if utility.IsErr(bindErr, "BindJSON Failed!") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server JSON bind failed",
		})
		return
	}
	service := jsondata.Service
	namespace := jsondata.Namespace
	var output bytes.Buffer
	if !utility.KubeDeleteService(service, namespace, &output) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delete service failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Service delete succeed",
		"data":    output.String(),
	})
	return

}

func DeleteDeploy(c *gin.Context) {
	type msg struct {
		Deploy    string `json:"deploy"`
		Namespace string `json:"namespace"`
	}
	jsondata := msg{}
	bindErr := c.BindJSON(&jsondata)
	if utility.IsErr(bindErr, "BindJSON Failed!") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server JSON bind failed",
		})
		return
	}
	deploy := jsondata.Deploy
	namespace := jsondata.Namespace
	var output bytes.Buffer
	if !utility.KubeDeleteDeploy(deploy, namespace, &output) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Delete deploy failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deploy delete succeed",
		"data":    output.String(),
	})
	return
}

func GetDockerImages(c *gin.Context) {
	var output bytes.Buffer
	if utility.DockerImages(&output) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Docker images run succeed",
			"data":    output.String(),
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Docker images run failed",
	})
	return
}

func GetTop(c *gin.Context) {
	type msg struct {
		Spec string `json:"spec"`
	}
	jsondata := msg{}
	bindErr := c.BindJSON(&jsondata)
	if utility.IsErr(bindErr, "BindJSON Failed!") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server JSON bind failed",
		})
		return
	}
	spec := jsondata.Spec
	var output bytes.Buffer
	if utility.KubeTop(spec, &output) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Kube top run succeed",
			"data":    output.String(),
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Kube top run failed",
	})
	return
}

func CarControl(c *gin.Context) {
	type msg struct {
		Control string `json:"control"`
	}
	jsondata := msg{}
	bindErr := c.BindJSON(&jsondata)
	if utility.IsErr(bindErr, "BindJSON Failed!") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server JSON bind failed",
		})
		return
	}
	buf := jsondata.Control
	if utility.SendCarControl(buf) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Car control send succeed",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Car control send failed",
	})
	return
}
