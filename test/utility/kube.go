package utility

import (
	"bytes"
	"os/exec"
)

const kubeLogDir string = "./log/kubeLog/log.txt"

func KubeApplyYaml(filename string, output *bytes.Buffer) bool {
	dir := "tmp/"
	file := dir + filename
	cmdline := "kubectl apply -f " + file
	cmd := exec.Command("/bin/bash", "-c", cmdline)
	cmd.Stdout = output
	cmdRunErr := cmd.Run()
	if !IsErr(cmdRunErr, "Kubectl Apply Yaml Failed!") {
		WriteLog(kubeLogDir, cmdline+"\n"+output.String())
		return true
	}
	return false
}

func KubeDeletService(service string, namespace string, output *bytes.Buffer) bool {
	cmdline := "kubectl delete service " + service + " " + "--namespace=" + namespace
	cmd := exec.Command("/bin/bash", "-c", cmdline)
	cmd.Stdout = output
	cmdRunErr := cmd.Run()
	if !IsErr(cmdRunErr, "Kubectl Delete Service Failed!") {
		WriteLog(kubeLogDir, cmdline+"\n"+output.String())
		return true
	}
	return false
}
