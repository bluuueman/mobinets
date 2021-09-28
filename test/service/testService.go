package service

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

//文件上传
func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("ERROR: upload file failed. ", err)
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
	if bindErr != nil {
		log.Println("BindJSON error! ", bindErr)
	}
	target := "tmp/writefile.txt"
	cmd := exec.Command("/bin/bash", "-c", "echo"+" "+jsondata.Message+" > "+target)
	cmdRunErr := cmd.Run()
	if cmdRunErr != nil {
		log.Println("ERROR: Command exec failed. ", cmdRunErr)
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