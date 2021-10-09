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
}
