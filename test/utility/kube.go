package utility

import (
	"bytes"
	"os/exec"
)

const kubeLogDir string = "./log/kubeLog/log.txt"

func KubeApplyYaml(filename string, output *bytes.Buffer) bool {
	dir := "tmp/"
	file := dir + filename
	cmd := exec.Command("/bin/bash", "-c", "kubectl apply -f "+file)
	cmd.Stdout = output
	cmdRunErr := cmd.Run()
	if !IsErr(cmdRunErr, "Kubectl Apply Yaml Failed!") {
		WriteLog(kubeLogDir, "kubectl apply -f "+file+"/n"+output.String())
		return true
	}
	return false
}
