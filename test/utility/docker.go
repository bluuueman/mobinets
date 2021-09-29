package utility

import (
	"os"
	"os/exec"
)
logDir := "test/tmp/dockerLog.txt"

func dockerPull(image string) *exec.Command{
	
	cmd := exec.Command("/bin/bash", "-c", "docker pull "+image)
	cmdStartErr := cmd.Start()
	if cmdStartErr != nil{
		log.Println(cmdStartErr)
		return nil
	}
	return cmd
}

func dockerPullResult(cmd *exec.Cmd) int{
	cmdWaitErr := cmd.Wait()
	if cmdWaitErr != nil{
		log.Println(cmdWaitErr)

		return 0
	}
	writeLog()
	return 1
	
}

func dockerImages() int{
	cmd := exec.Command("/bin/bash", "-c", "docker images")

}