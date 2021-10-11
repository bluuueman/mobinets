package utility

import (
	"bytes"
	"os/exec"
)

const dockerLogDir string = "./log/dockerLog/log.txt"

func DockerPull(image string, output *bytes.Buffer) *exec.Cmd {

	cmd := exec.Command("/bin/bash", "-c", "docker pull "+image)
	cmd.Stdout = output
	err := cmd.Start()
	if IsErr(err, "Command Start Failed!") {
		return nil
	}
	return cmd
}

func DockerPullResult(cmd *exec.Cmd) bool {
	err := cmd.Wait()
	if IsErr(err, "Command Wait Failed!") {
		return false
	}
	return true
}

func DockerImages(output *bytes.Buffer) bool {
	cmd := exec.Command("/bin/bash", "-c", "docker images")
	cmd.Stdout = output
	err := cmd.Run()
	if !IsErr(err, "Command Run Failed!") {
		WriteLog(dockerLogDir, "CMD: docker images\n"+output.String())
		return true
	}
	return false

}
